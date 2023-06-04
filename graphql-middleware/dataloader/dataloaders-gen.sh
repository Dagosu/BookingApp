#!/bin/bash
set -uexo pipefail
cd "$(dirname "${BASH_SOURCE[0]}")"

print () { printf "\033[1;33m$1\033[0m\n"; }

gen_dataloader() {
  LOADER_NAME=$1
  MODEL_NAME=$2
  GENERATED_FILE="$3_gen.go"

  if [[ ! -f "$GENERATED_FILE" || $GENERATED_FILE -ot "dataloaders.go" ]]; then
      go run \
        github.com/vektah/dataloaden \
        $LOADER_NAME \
        "gitlab.com/airportlabs/graphql-common/model.Transport" \
        "*gitlab.com/airportlabs/graphql-common/model.$MODEL_NAME"
      echo " - [✓] $LOADER_NAME"
  fi
}

print "Generating dataloader files..."

gen_dataloader FittingLoader Fitting fittingloader
gen_dataloader FieldDescriptorLoader FieldDescriptor fielddescriptorloader
gen_dataloader RoleLoader Role roleloader
