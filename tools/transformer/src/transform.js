const {
  applyAllOfTransformations,
  applyOneOfTransformations,
  applyModelNameTransformations,
  applyDiscriminatorTransformations,
  applyDigestTransformations,
} = require("./transformations");
const { getAPI, saveAPI } = require("./engine/apifile");
const {  } = require("./transformations/digest");

const log = require("simple-node-logger").createSimpleLogger();

// Override default logger
global.console = log;
log.setLevel(process.env.XGEN_LOGGING_LEVEL || "warn");

let { doc, apiFileLocation } = getAPI(process.argv.slice(2));

doc = applyDiscriminatorTransformations(doc);
doc = applyOneOfTransformations(doc);
doc = applyAllOfTransformations(doc);
doc = applyModelNameTransformations(doc, "ApiAtlas", "View");
doc = applyDigestTransformations(doc);


// CLI generator

// Generate CLI per path. 
const cliSamplePath = "/api/atlas/v2/orgs";
const value = doc.paths[cliSamplePath]
doc.paths = {"/api/atlas/v2/orgs": value}

docs.paths["/api/atlas/v2/orgs"].paramters

saveAPI(doc, apiFileLocation);
