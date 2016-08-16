'use strict'
const path = require('path')
const gcloud = require('gcloud')

function upload(config, bucketName, fileName) {
  const gcs = gcloud.storage(config)
  const bucket = gcs.bucket(bucketName)

  bucket.upload(path.join(GOOSE_PATH, fileName), (err) => {
    if (err) {
      console.log(`Failed to insert ${fileName} because of `, err)
      throw err
    } else {
      console.log('Successfully uploaded ', fileName)
    }
  })
}

const args = process.argv
const client_email = args[2]
const private_key = args[3]

const GOOSE_PATH = path.join(__dirname, '..', 'cmd', 'goose')

const config = {
  projectId: 'help-1272'
  , credentials: {
    client_email
    , private_key
  }
}
const bucketName = 'helpdotcom-binaries'

upload(config, bucketName, 'goose')
