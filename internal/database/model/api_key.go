package database_model

const (
	GENERAL = "GENERAL"
	SUPER_ADMIN = "SUPER_ADMIN"
)

const (
	ACTIVE = 1
	INACTIVE = 0
)

type ApiKey struct {
	_id          uint64 `json:"id"`
	key          string `json:"key"`
	permissions  string `json:"permissions"`
	createdAt    uint64 `json:"createdAt"`
	updatedAt    uint64 `json:"updatedAt"`
}

/*
export async function createSuperAdminApiKey() {
  // Check if there are any documents in the collection
  const count = ApiKeyModel.All.length;

  // If no documents exist, add a default document
  if (count === 0) {
    const defaultDocument: ApiKey = {
      _id: 0,
      key: `${superAdminApiKey}`,
      permissions: Permission.SUPER_ADMIN,
      status: Status.ACTIVE,
      createdAt: new Date().toString(),
      updatedAt: new Date().toString(),
    };

    // Save the default document
    const result = ApiKeyModel.InsertOne(defaultDocument);
    if (!(await result).success) throw new InternalError((await result).error);
  }
}
*/
