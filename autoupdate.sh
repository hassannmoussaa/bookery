set -ex

PROJECT_ID="bookery"
PROJECT_ROOT_DIR="/bookery"
APP_DIR="$PROJECT_ROOT_DIR/app"
LOGS_DIR="/bookery/logs"
PROJECT_URL="https://storage.googleapis.com/bookery-binary/bookery.tar"
BINARY_FILE_PATH="$APP_DIR/$PROJECT_ID"
ENVIRONMENT="production"
DOMAIN_NAME="bookery"

sudo apt-get update -yq && sudo apt-get upgrade -yq
apt-get autoremove -yq

# Get the application tar from the GCS bucket.
if [ ! -d $PROJECT_ROOT_DIR ]; then
mkdir $PROJECT_ROOT_DIR
fi
if [ -f "$PROJECT_ROOT_DIR/$PROJECT_ID.tar" ]; then
rm "$PROJECT_ROOT_DIR/$PROJECT_ID.tar"
fi
curl -H 'Cache-Control: no-cache' -o "$PROJECT_ROOT_DIR/$PROJECT_ID.tar" $PROJECT_URL
if [ -d $APP_DIR ]; then
   if [ -L $APP_DIR ]; then 
      rm $APP_DIR
   else 
      rm -rf $APP_DIR
   fi
fi
mkdir $APP_DIR
if [ ! -d $LOGS_DIR ]; then
mkdir $LOGS_DIR
fi
tar -x -f "$PROJECT_ROOT_DIR/$PROJECT_ID.tar" -C $APP_DIR
chmod +x $BINARY_FILE_PATH

sed -i "s/domain-name/$DOMAIN_NAME/g" "$APP_DIR/$ENVIRONMENT.json"

#Reboot server
sv stop $PROJECT_ID
sv start $PROJECT_ID
# Application should now be running under runit