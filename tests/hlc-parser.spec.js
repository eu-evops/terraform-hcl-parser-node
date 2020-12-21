const { expect } = require('chai');
const util = require('util');
const { readFileSync } = require('fs');
const { join } = require('path');

const parser = require('../dist');

describe('hlc-parser', () => {
  it('should parse simple terraform HLC', () => {
    const moduleObject = parser.parse('module test {}')

    expect(moduleObject).to.have.property('module_calls').to.have.property('test')
    expect(moduleObject).to.have.property('path').to.eq('<virtual>')
  })

  it('should parse full terraform HLC', () => {
    const moduleObject = parser.parse(readFileSync(join(__dirname, './data/full/main.tf')))

    expect(moduleObject).to.have.property('path').to.eq('<virtual>')
    expect(moduleObject).to.have.property('module_calls').to.have.property('storage')
    expect(moduleObject.module_calls.storage).to.have.property('source').to.eq('./modules/storage')
    expect(moduleObject).to.have.property('variables').to.have.property('test')
    expect(moduleObject).to.have.property('outputs').to.have.property('test')
    expect(moduleObject).to.have.property('data_resources').to.have.property('data.test_data.test')
    expect(moduleObject).to.have.property('managed_resources').to.have.property('test_resource.test')
  })
})