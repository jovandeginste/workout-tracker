# GORM Data Types

## JSON

sqlite, mysql, postgres supported

```go
import "gorm.io/datatypes"

type UserWithJSON struct {
	gorm.Model
	Name       string
	Attributes datatypes.JSON
}

DB.Create(&UserWithJSON{
	Name:       "json-1",
	Attributes: datatypes.JSON([]byte(`{"name": "jinzhu", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
}

// Check JSON has keys
datatypes.JSONQuery("attributes").HasKey(value, keys...)

db.Find(&user, datatypes.JSONQuery("attributes").HasKey("role"))
db.Find(&user, datatypes.JSONQuery("attributes").HasKey("orgs", "orga"))
// MySQL
// SELECT * FROM `users` WHERE JSON_EXTRACT(`attributes`, '$.role') IS NOT NULL
// SELECT * FROM `users` WHERE JSON_EXTRACT(`attributes`, '$.orgs.orga') IS NOT NULL

// PostgreSQL
// SELECT * FROM "user" WHERE "attributes"::jsonb ? 'role'
// SELECT * FROM "user" WHERE "attributes"::jsonb -> 'orgs' ? 'orga'


// Check JSON extract value from keys equal to value
datatypes.JSONQuery("attributes").Equals(value, keys...)

DB.First(&user, datatypes.JSONQuery("attributes").Equals("jinzhu", "name"))
DB.First(&user, datatypes.JSONQuery("attributes").Equals("orgb", "orgs", "orgb"))
// MySQL
// SELECT * FROM `user` WHERE JSON_EXTRACT(`attributes`, '$.name') = "jinzhu"
// SELECT * FROM `user` WHERE JSON_EXTRACT(`attributes`, '$.orgs.orgb') = "orgb"

// PostgreSQL
// SELECT * FROM "user" WHERE json_extract_path_text("attributes"::json,'name') = 'jinzhu'
// SELECT * FROM "user" WHERE json_extract_path_text("attributes"::json,'orgs','orgb') = 'orgb'
```

NOTE: SQlite need to build with `json1` tag, e.g: `go build --tags json1`, refer https://github.com/mattn/go-sqlite3#usage

## Date

```go
import "gorm.io/datatypes"

type UserWithDate struct {
	gorm.Model
	Name string
	Date datatypes.Date
}

user := UserWithDate{Name: "jinzhu", Date: datatypes.Date(time.Now())}
DB.Create(&user)
// INSERT INTO `user_with_dates` (`name`,`date`) VALUES ("jinzhu","2020-07-17 00:00:00")

DB.First(&result, "name = ? AND date = ?", "jinzhu", datatypes.Date(curTime))
// SELECT * FROM user_with_dates WHERE name = "jinzhu" AND date = "2020-07-17 00:00:00" ORDER BY `user_with_dates`.`id` LIMIT 1
```

## Time

MySQL, PostgreSQL, SQLite, SQLServer are supported.

Time with nanoseconds is supported for some databases which support for time with fractional second scale.

```go
import "gorm.io/datatypes"

type UserWithTime struct {
    gorm.Model
    Name string
    Time datatypes.Time
}

user := UserWithTime{Name: "jinzhu", Time: datatypes.NewTime(1, 2, 3, 0)}
DB.Create(&user)
// INSERT INTO `user_with_times` (`name`,`time`) VALUES ("jinzhu","01:02:03")

DB.First(&result, "name = ? AND time = ?", "jinzhu", datatypes.NewTime(1, 2, 3, 0))
// SELECT * FROM user_with_times WHERE name = "jinzhu" AND time = "01:02:03" ORDER BY `user_with_times`.`id` LIMIT 1
```

NOTE: If the current using database is SQLite, the field column type is defined as `TEXT` type
when GORM AutoMigrate because SQLite doesn't have time type.

## JSON_SET

sqlite, mysql, postgres supported

```go
import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type UserWithJSON struct {
	gorm.Model
	Name       string
	Attributes datatypes.JSON
}

DB.Create(&UserWithJSON{
	Name:       "json-1",
	Attributes: datatypes.JSON([]byte(`{"name": "json-1", "age": 18, "tags": ["tag1", "tag2"], "orgs": {"orga": "orga"}}`)),
})

type User struct {
	Name string
	Age  int
}

friend := User{
	Name: "Bob",
	Age:  21,
}

// Set fields of JSON column
datatypes.JSONSet("attributes").Set("age", 20).Set("tags[0]", "tag2").Set("orgs.orga", "orgb")

DB.Model(&UserWithJSON{}).Where("name = ?", "json-1").UpdateColumn("attributes", datatypes.JSONSet("attributes").Set("age", 20).Set("tags[0]", "tag3").Set("orgs.orga", "orgb"))
DB.Model(&UserWithJSON{}).Where("name = ?", "json-1").UpdateColumn("attributes", datatypes.JSONSet("attributes").Set("phones", []string{"10085", "10086"}))
DB.Model(&UserWithJSON{}).Where("name = ?", "json-1").UpdateColumn("attributes", datatypes.JSONSet("attributes").Set("phones", gorm.Expr("CAST(? AS JSON)", `["10085", "10086"]`)))
DB.Model(&UserWithJSON{}).Where("name = ?", "json-1").UpdateColumn("attributes", datatypes.JSONSet("attributes").Set("friend", friend))
// MySQL
// UPDATE `user_with_jsons` SET `attributes` = JSON_SET(`attributes`, '$.tags[0]', 'tag3', '$.orgs.orga', 'orgb', '$.age', 20) WHERE name = 'json-1'
// UPDATE `user_with_jsons` SET `attributes` = JSON_SET(`attributes`, '$.phones', CAST('["10085", "10086"]' AS JSON)) WHERE name = 'json-1'
// UPDATE `user_with_jsons` SET `attributes` = JSON_SET(`attributes`, '$.phones', CAST('["10085", "10086"]' AS JSON)) WHERE name = 'json-1'
// UPDATE `user_with_jsons` SET `attributes` = JSON_SET(`attributes`, '$.friend', CAST('{"Name": "Bob", "Age": 21}' AS JSON)) WHERE name = 'json-1'
```
NOTE: MariaDB does not support CAST(? AS JSON).

NOTE: Path in PostgreSQL is different.

```go
// Set fields of JSON column
datatypes.JSONSet("attributes").Set("{age}", 20).Set("{tags, 0}", "tag2").Set("{orgs, orga}", "orgb")

DB.Model(&UserWithJSON{}).Where("name = ?", "json-1").UpdateColumn("attributes", datatypes.JSONSet("attributes").Set("{age}", 20).Set("{tags, 0}", "tag2").Set("{orgs, orga}", "orgb"))
DB.Model(&UserWithJSON{}).Where("name = ?", "json-1").UpdateColumn("attributes", datatypes.JSONSet("attributes").Set("{phones}", []string{"10085", "10086"}))
DB.Model(&UserWithJSON{}).Where("name = ?", "json-1").UpdateColumn("attributes", datatypes.JSONSet("attributes").Set("{phones}", gorm.Expr("?::jsonb", `["10085", "10086"]`)))
DB.Model(&UserWithJSON{}).Where("name = ?", "json-1").UpdateColumn("attributes", datatypes.JSONSet("attributes").Set("{friend}", friend))
// PostgreSQL
// UPDATE "user_with_jsons" SET "attributes" = JSONB_SET(JSONB_SET(JSONB_SET("attributes", '{age}', '20'), '{tags, 0}', '"tag2"'), '{orgs, orga}', '"orgb"') WHERE name = 'json-1'
// UPDATE "user_with_jsons" SET "attributes" = JSONB_SET("attributes", '{phones}', '["10085","10086"]') WHERE name = 'json-1'
// UPDATE "user_with_jsons" SET "attributes" = JSONB_SET("attributes", '{phones}', '["10085","10086"]'::jsonb) WHERE name = 'json-1'
// UPDATE "user_with_jsons" SET "attributes" = JSONB_SET("attributes", '{friend}', '{"Name": "Bob", "Age": 21}') WHERE name = 'json-1'
```

## JSONType[T]

sqlite, mysql, postgres supported

```go
import "gorm.io/datatypes"

type Attribute struct {
	Sex   int
	Age   int
	Orgs  map[string]string
	Tags  []string
	Admin bool
	Role  string
}

type UserWithJSON struct {
	gorm.Model
	Name       string
	Attributes datatypes.JSONType[Attribute]
}

var user = UserWithJSON{
	Name: "hello",
	Attributes: datatypes.NewJSONType(Attribute{
        Age:  18,
        Sex:  1,
        Orgs: map[string]string{"orga": "orga"},
        Tags: []string{"tag1", "tag2", "tag3"},
    }),
}

// Create
DB.Create(&user)

// First
var result UserWithJSON
DB.First(&result, user.ID)

// Update
jsonMap = UserWithJSON{
	Attributes: datatypes.NewJSONType(Attribute{
        Age:  18,
        Sex:  1,
        Orgs: map[string]string{"orga": "orga"},
        Tags: []string{"tag1", "tag2", "tag3"},
    }),
}

DB.Model(&user).Updates(jsonMap)
```

NOTE: it's not support json query

## JSONSlice[T]

sqlite, mysql, postgres supported

```go
import "gorm.io/datatypes"

type Tag struct {
	Name  string
	Score float64
}

type UserWithJSON struct {
	gorm.Model
	Name       string
	Tags       datatypes.JSONSlice[Tag]
}

var tags = []Tag{{Name: "tag1", Score: 0.1}, {Name: "tag2", Score: 0.2}}
var user = UserWithJSON{
	Name: "hello",
	Tags: datatypes.NewJSONSlice(tags),
}

// Create
DB.Create(&user)

// First
var result UserWithJSON
DB.First(&result, user.ID)

// Update
var tags2 = []Tag{{Name: "tag3", Score: 10.1}, {Name: "tag4", Score: 10.2}}
jsonMap = UserWithJSON{
	Tags: datatypes.NewJSONSlice(tags2),
}

DB.Model(&user).Updates(jsonMap)
```

NOTE: it's not support json query and `db.Pluck` method

## JSONArray

mysql supported

```go
import "gorm.io/datatypes"

type Param struct {
    ID          int
    Letters     string
    Config      datatypes.JSON
}

//Create
DB.Create(&Param{
    Letters: "JSONArray-1",
    Config:      datatypes.JSON("[\"a\", \"b\"]"),
})

DB.Create(&Param{
    Letters: "JSONArray-2",
    Config:      datatypes.JSON("[\"a\", \"c\"]"),
})

//Query
var retMultiple []Param
DB.Where(datatypes.JSONArrayQuery("config").Contains("c")).Find(&retMultiple)
}
```

## UUID

MySQL, PostgreSQL, SQLServer and SQLite are supported.

```go
import "gorm.io/datatypes"

type UserWithUUID struct {
    gorm.Model
    Name string
    UserUUID datatypes.UUID
}

// Generate a new random UUID (version 4).
userUUID := datatypes.NewUUIDv4()

user := UserWithUUID{Name: "jinzhu", UserUUID: userUUID}
DB.Create(&user)
// INSERT INTO `user_with_uuids` (`name`,`user_uuid`) VALUES ("jinzhu","ca95a578-816c-4812-babd-a7602b042460")

var result UserWithUUID
DB.First(&result, "name = ? AND user_uuid = ?", "jinzhu", userUUID)
// SELECT * FROM user_with_uuids WHERE name = "jinzhu" AND user_uuid = "ca95a578-816c-4812-babd-a7602b042460" ORDER BY `user_with_uuids`.`id` LIMIT 1

// Use the datatype's Equals() to compare the UUIDs.
if userCreate.UserUUID.Equals(userFound.UserUUID) {
	fmt.Println("User UUIDs match as expected.")
} else {
	fmt.Println("User UUIDs do not match. Something is wrong.")
}

// Use the datatype's String() function to get the UUID as a string type.
fmt.Printf("User UUID is %s", userFound.UserUUID.String())

// Check the UUID value with datatype's IsNil() and IsEmpty() functions.
if userFound.UserUUID.IsNil() {
	fmt.Println("User UUID is a nil UUID (i.e. all bits are zero)")
}
if userFound.UserUUID.IsEmpty() {
	fmt.Println(
		"User UUID is empty (i.e. either a nil UUID or a zero length string)",
	)
}
```
