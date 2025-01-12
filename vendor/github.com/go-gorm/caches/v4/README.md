# Gorm Caches

Gorm Caches plugin using database request reductions (easer), and response caching mechanism provide you an easy way to optimize database performance.

## Features

- Database request reduction. If three identical requests are running at the same time, only the first one is going to be executed, and its response will be returned for all.
- Database response caching. By implementing the Cacher interface, you can easily setup a caching mechanism for your database queries.
- Supports all databases that are supported by gorm itself.

## Install

```bash
go get -u github.com/go-gorm/caches/v4
```

## Usage

Configure the `easer`, and the `cacher`, and then load the plugin to gorm.

```go
package main

import (
	"fmt"
	"sync"

	"github.com/go-gorm/caches/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(
		mysql.Open("DATABASE_DSN"),
		&gorm.Config{},
	)
	cachesPlugin := &caches.Caches{Conf: &caches.Config{
		Easer: true,
		Cacher: &yourCacherImplementation{},
	}}
	_ = db.Use(cachesPlugin)
}
```

## Easer Example

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-gorm/caches/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserRoleModel struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type UserModel struct {
	gorm.Model
	Name   string
	RoleId uint
	Role   *UserRoleModel `gorm:"foreignKey:role_id;references:id"`
}

func main() {
	db, _ := gorm.Open(
		mysql.Open("DATABASE_DSN"),
		&gorm.Config{},
	)

	cachesPlugin := &caches.Caches{Conf: &caches.Config{
		Easer: true,
	}}

	_ = db.Use(cachesPlugin)

	_ = db.AutoMigrate(&UserRoleModel{})

	_ = db.AutoMigrate(&UserModel{})

	adminRole := &UserRoleModel{
		Name: "Admin",
	}
	db.FirstOrCreate(adminRole, "Name = ?", "Admin")

	guestRole := &UserRoleModel{
		Name: "Guest",
	}
	db.FirstOrCreate(guestRole, "Name = ?", "Guest")

	db.Save(&UserModel{
		Name: "ktsivkov",
		Role: adminRole,
	})
	db.Save(&UserModel{
		Name: "anonymous",
		Role: guestRole,
	})

	var (
		q1Users []UserModel
		q2Users []UserModel
	)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		db.Model(&UserModel{}).Joins("Role").Find(&q1Users, "Role.Name = ? AND Sleep(1) = false", "Admin")
		wg.Done()
	}()
	go func() {
		time.Sleep(500 * time.Millisecond)
		db.Model(&UserModel{}).Joins("Role").Find(&q2Users, "Role.Name = ? AND Sleep(1) = false", "Admin")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println(fmt.Sprintf("%+v", q1Users))
	fmt.Println(fmt.Sprintf("%+v", q2Users))
}
```

## Cacher Example (Redis)

```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-gorm/caches/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRoleModel struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type UserModel struct {
	gorm.Model
	Name   string
	RoleId uint
	Role   *UserRoleModel `gorm:"foreignKey:role_id;references:id"`
}

type redisCacher struct {
	rdb *redis.Client
}

func (c *redisCacher) Get(ctx context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
	res, err := c.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if err := q.Unmarshal([]byte(res)); err != nil {
		return nil, err
	}

	return q, nil
}

func (c *redisCacher) Store(ctx context.Context, key string, val *caches.Query[any]) error {
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	c.rdb.Set(ctx, key, res, 300*time.Second) // Set proper cache time
	return nil
}

func (c *redisCacher) Invalidate(ctx context.Context) error {
	var (
		cursor uint64
		keys   []string
	)
	for {
		var (
			k   []string
			err error
		)
		k, cursor, err = c.rdb.Scan(ctx, cursor, fmt.Sprintf("%s*", caches.IdentifierPrefix), 0).Result()
		if err != nil {
			return err
		}
		keys = append(keys, k...)
		if cursor == 0 {
			break
		}
	}

	if len(keys) > 0 {
		if _, err := c.rdb.Del(ctx, keys...).Result(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		AllowGlobalUpdate: true,
	})

	cachesPlugin := &caches.Caches{Conf: &caches.Config{
		Cacher: &redisCacher{
			rdb: redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				Password: "",
				DB:       0,
			}),
		},
	}}

	_ = db.Use(cachesPlugin)

	_ = db.AutoMigrate(&UserRoleModel{})
	_ = db.AutoMigrate(&UserModel{})

	db.Delete(&UserRoleModel{})
	db.Delete(&UserModel{})

	adminRole := &UserRoleModel{
		Name: "Admin",
	}
	db.Save(adminRole)

	guestRole := &UserRoleModel{
		Name: "Guest",
	}
	db.Save(guestRole)

	db.Save(&UserModel{
		Name: "ktsivkov",
		Role: adminRole,
	})

	db.Save(&UserModel{
		Name: "anonymous",
		Role: guestRole,
	})

	q1User := &UserModel{}
	db.WithContext(context.Background()).Find(q1User, "Name = ?", "ktsivkov")
	q2User := &UserModel{}
	db.WithContext(context.Background()).Find(q2User, "Name = ?", "ktsivkov")

	fmt.Println(fmt.Sprintf("%+v", q1User))
	fmt.Println(fmt.Sprintf("%+v", q2User))
}
```

## Cacher Example (Memory)

```go
package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-gorm/caches/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserRoleModel struct {
	gorm.Model
	Name string `gorm:"unique"`
}

type UserModel struct {
	gorm.Model
	Name   string
	RoleId uint
	Role   *UserRoleModel `gorm:"foreignKey:role_id;references:id"`
}

type memoryCacher struct {
	store *sync.Map
}

func (c *memoryCacher) init() {
	if c.store == nil {
		c.store = &sync.Map{}
	}
}

func (c *memoryCacher) Get(ctx context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
	c.init()
	val, ok := c.store.Load(key)
	if !ok {
		return nil, nil
	}

	if err := q.Unmarshal(val.([]byte)); err != nil {
		return nil, err
	}

	return q, nil
}

func (c *memoryCacher) Store(ctx context.Context, key string, val *caches.Query[any]) error {
	c.init()
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	c.store.Store(key, res)
	return nil
}

func (c *memoryCacher) Invalidate(ctx context.Context) error {
	c.store = &sync.Map{}
	return nil
}

func main() {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		AllowGlobalUpdate: true,
	})

	cachesPlugin := &caches.Caches{Conf: &caches.Config{
		Cacher: &memoryCacher{},
	}}

	_ = db.Use(cachesPlugin)

	_ = db.AutoMigrate(&UserRoleModel{})
	_ = db.AutoMigrate(&UserModel{})

	db.Delete(&UserRoleModel{})
	db.Delete(&UserModel{})

	adminRole := &UserRoleModel{
		Name: "Admin",
	}
	db.Save(adminRole)

	guestRole := &UserRoleModel{
		Name: "Guest",
	}
	db.Save(guestRole)

	db.Save(&UserModel{
		Name: "ktsivkov",
		Role: adminRole,
	})

	db.Save(&UserModel{
		Name: "anonymous",
		Role: guestRole,
	})

	q1User := &UserModel{}
	db.WithContext(context.Background()).Find(q1User, "Name = ?", "ktsivkov")
	q2User := &UserModel{}
	db.WithContext(context.Background()).Find(q2User, "Name = ?", "ktsivkov")

	fmt.Println(fmt.Sprintf("%+v", q1User))
	fmt.Println(fmt.Sprintf("%+v", q2User))
}
```

## License

MIT license.

## Easer
The easer is an adjusted version of the [ServantGo](https://github.com/ktsivkov/servantgo) library to fit the needs of this plugin.
