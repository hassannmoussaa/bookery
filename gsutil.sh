BUCKET_NAME="bookery-app-public"
STATIC_FILES_DIR=/Users/nehme/go/src/github.com/hassannmoussaa/bookery/static/dist/*

gsutil -m cp -r -z js,css $STATIC_FILES_DIR  gs://$BUCKET_NAME/static/dist
#gsutil acl ch -u AllUsers:R gs://$BUCKET_NAME/**

gsutil -m setmeta -h "Content-Type:application/javascript" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.js

gsutil -m setmeta -h "Content-Type:text/css" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.css

gsutil -m setmeta -h "Content-Type:image/png" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.png

gsutil -m setmeta -h "Content-Type:image/jpeg" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.jpg

gsutil -m setmeta -h "Content-Type:image/jpeg" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.jpeg

gsutil -m setmeta -h "Content-Type:image/gif" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.gif

gsutil -m setmeta -h "Content-Type:application/pdf" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.pdf

gsutil -m setmeta -h "Content-Type:image/svg+xml" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.svg

gsutil -m setmeta -h "Content-Type:application/x-font-ttf" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.ttf

gsutil -m setmeta -h "Content-Type:application/font-woff" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.woff

gsutil -m setmeta -h "Content-Type:application/vnd.ms-fontobject" \
  -h "Cache-Control:public, max-age=15552000" \
  -h "Content-Disposition" gs://$BUCKET_NAME/static/**/*.eot
