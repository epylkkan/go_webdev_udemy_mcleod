
https://cloud.google.com/sql/docs/mysql/quickstart-proxy-test


Install CloudSQL Auth proxy 
https://cloud.google.com/sql/docs/mysql/quickstart-proxy-test
chmod +x cloud_sql_proxy

Start CloudAQL Auth proxy

./cloud_sql_proxy -instances=heroic-oarlock-340615:europe-north1:gotraining=tcp:3306
mysql -u root -p --host 127.0.0.1 --port 3306


install mysql (community edition)
install workbench
create mysql db on gcp

conenct from GCP cloud shell: gcloud sql connect gotraining --user=root --quiet
