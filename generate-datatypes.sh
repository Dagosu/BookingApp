#!/bin/bash

# Run protoc
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:. --go-grpc_opt=paths=source_relative flight.proto operation.proto field_descriptor.proto field_type.proto user.proto

# Find all generated pb.go files
for file in *.pb.go
do
  # Add bson tags after json tags
  sed -i '' -e 's/json:"\([^"]*\)"/json:"\1" bson:"\1"/g' "$file"

  # Handle the "id" field specifically
 sed -i '' -E 's/json:"id,omitempty"/json:"_id,omitempty"/' "$file"

  # Handle the "id" field specifically again after bson tags have been added
  sed -i '' -E 's/bson:"id,omitempty"/bson:"_id,omitempty"/' "$file"
done
