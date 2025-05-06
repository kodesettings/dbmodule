//go:build js && wasm
package database_model

const (
	PERMISSION_GENERAL = "GENERAL"
	PERMISSION_SUPER_ADMIN = "SUPER_ADMIN"
)

const (
	STATUS_ACTIVE = 1
	STATUS_INACTIVE = 0
)

type ApiKey struct {
	Id           uint64 `json:"id"`
	Key          string `json:"key"`
	Permissions  string `json:"permissions"`
	Status       uint16 `json:"status"`
	CreatedAt    uint64 `json:"createdAt"`
	UpdatedAt    uint64 `json:"updatedAt"`
}

func CreateSuperAdminApiKey() {}

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
