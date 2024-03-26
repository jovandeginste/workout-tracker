package units

import "errors"

var (
	// Shorthand for pre-defined unit systems
	BI  = UnitOptionSystem("imperial")
	SI  = UnitOptionSystem("metric")
	US  = UnitOptionSystem("us")
	IEC = UnitOptionSystem("iec")

	unitMap = make(map[string]Unit)
)

type Unit struct {
	Name     string
	Symbol   string
	Quantity string
	plural   string // either "none", "auto", or a specific plural name
	aliases  []string
	system   string
}

// NewUnit registers a new Unit within the package, returning the newly created Unit
func NewUnit(name, symbol string, opts ...UnitOption) Unit {
	if _, ok := unitMap[name]; ok {
		panic(errors.New("duplicate unit name: " + name))
	}

	u := Unit{
		Name:   name,
		Symbol: symbol,
		plural: "auto",
	}

	for _, opt := range opts {
		u = opt(u)
	}

	unitMap[name] = u
	return u
}

// Returns all names and symbols this unit may be referred to
func (u Unit) Names() []string {
	names := []string{u.Name}
	if u.Symbol != "" {
		names = append(names, u.Symbol)
	}
	if u.plural != "none" && u.plural != "auto" {
		names = append(names, u.PluralName())
	}
	return append(names, u.aliases...)
}

// Return the system of units this Unit belongs to, if any
func (u Unit) System() string { return u.system }

// Return the plural name for this unit
func (u Unit) PluralName() string {
	switch u.plural {
	case "none":
		return u.Name
	case "auto":
		return u.Name + "s"
	default: // custom plural name
		return u.plural
	}
}

// Option that may be passed to NewUnit
type UnitOption func(Unit) Unit

// Either "none", "auto", or a custom plural unit name
// "none" - labels will use the unmodified unit name in a plural context
// "auto" - labels for this unit will be created with a plural suffix when appropriate (default)
func UnitOptionPlural(s string) UnitOption {
	return func(u Unit) Unit {
		u.plural = s
		return u
	}
}

// Additional names, spellings, or symbols that this unit may be referred to as
func UnitOptionAliases(a ...string) UnitOption {
	return func(u Unit) Unit {
		u.aliases = append(u.aliases, a...)
		return u
	}
}

// Set a system of units for which this Unit belongs
func UnitOptionSystem(s string) UnitOption {
	return func(u Unit) Unit {
		u.system = s
		return u
	}
}

// Set a quantity label for which this Unit belongs
func UnitOptionQuantity(s string) UnitOption {
	return func(u Unit) Unit {
		u.Quantity = s
		return u
	}
}

type UnitList []Unit

// UnitList implements sort.Interface
func (a UnitList) Len() int      { return len(a) }
func (a UnitList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a UnitList) Less(i, j int) bool {
	if a[i].Quantity != a[j].Quantity {
		return a[i].Quantity < a[j].Quantity
	}
	return a[i].Name < a[j].Name
}
