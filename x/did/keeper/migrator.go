package keeper

type Migrator struct {
	Keeper Keeper
}

func NewMigrator(k Keeper) Migrator {
	return Migrator{
		Keeper: k,
	}
}
