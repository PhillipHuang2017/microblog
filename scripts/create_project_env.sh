cd ..
mkdir -p project_env/mysql/init && mkdir project_env/mysql/conf && cp ./service/rpc/user/sql/user.sql project_env/mysql/init/user.sql
cp etc/mysql.cnf project_env/mysql/conf/mysql.cnf

mkdir project_env/redis

mkdir project_env/user-api && cp ./service/api/user/etc/user-api.yaml project_env/user-api/user-api.yaml
mkdir project_env/user-rpc && cp ./service/rpc/user/etc/user.yaml project_env/user-rpc/user.yaml
cd -