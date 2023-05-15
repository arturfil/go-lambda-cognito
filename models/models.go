package models

type SecretRDSJson struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	DBClusterIdentifier string `json:"db_cluster_identifier"`
}

type SignUp struct {
	UserEmail string `json:"user_email"`
	UserUUID  string `json:"user_uuid"`
}
