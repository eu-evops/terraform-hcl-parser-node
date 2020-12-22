![Go](https://github.com/eu-evops/terraform-hcl-terraform-parser-node/workflows/Go/badge.svg)

# Terraform HCL Parser for Node.js

Using gopherjs to transpile Hashicorp's HCL into javascript.

Inspired by [anhldbk/hcl-parser](https://github.com/anhldbk/hcl-parser)

# Usage

```javascript
const fs = require('fs');
const parser = require('@evops/hcl-terraform-parser');

const content = fs.readFileSync('main.tf');
const hclFile = parser.parse(content);

```

`hclFile` is a JSON object with following structure:

```javascript
{
  path: '<virtual>',
  variables: {
    test: {
      name: 'test',
      default: 'Default value',
      required: false,
      pos: [Object]
    }
  },
  outputs: { test: { name: 'test', pos: [Object] } },
  required_providers: { test: { version_constraints: [Array] } },
  provider_configs: { test: { name: 'test' } },
  managed_resources: {
    'test_resource.test': {
      mode: 'managed',
      type: 'test_resource',
      name: 'test',
      provider: [Object],
      pos: [Object]
    }
  },
  data_resources: {
    'data.test_data.test': {
      mode: 'data',
      type: 'test_data',
      name: 'test',
      provider: [Object],
      pos: [Object]
    }
  },
  module_calls: {
    storage: { name: 'storage', source: './modules/storage', pos: [Object] }
  }
}
```


Credits:
* Fabian Ponce [FabianPonce](https://github.com/FabianPonce) for recommending `terraform-config-inspect`
