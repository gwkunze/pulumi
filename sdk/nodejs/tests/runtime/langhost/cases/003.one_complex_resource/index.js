// This tests the creation of a resource that contains "simple" input and output propeprties.
// In particular, there aren't any fancy dataflow linked properties.

let assert = require("assert");
let fabric = require("../../../../../");

class MyResource extends fabric.Resource {
    constructor(name) {
        super("test:index:MyResource", name, {
            // First a few basic properties that are resolved to values.
            "inpropB1": false,
            "inpropB2": true,
            "inpropN": 42,
            "inpropS": "a string",
            "inpropA": [ true, 99, "what a great property" ],
            "inpropO": {
                b1: false,
                b2: true,
                n: 42,
                s: "another string",
                a: [ 66, false, "strings galore" ],
                o: { z: "x" },
            },

            // Next some properties that are completely unresolved (outputs).
            "outprop1": undefined,
            "outprop2": undefined,
        });
    }
}

let res = new MyResource("testResource1");
res.id.mapValue(id => {
    console.log(`ID: ${id}`);
    assert.equal(id, "testResource1");
});
res.urn.mapValue(urn => {
    console.log(`URN: ${urn}`);
    assert.equal(urn, "test:index:MyResource::testResource1");
});
res.outprop1.mapValue(prop => {
    console.log(`OutProp1: ${prop}`);
    assert.equal(prop, "output properties ftw");
});
res.outprop2.mapValue(prop => {
    console.log(`OutProp2: ${prop}`);
    assert.equal(prop, 998.6);
});
