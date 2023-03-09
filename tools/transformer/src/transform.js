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
// This transformation does filter entire api to single path
const paths = Object.keys(doc.paths).filter(path => path == "/api/atlas/v2/groups/{groupId}/limits");
console.warn("path", paths);

doc.paths = paths.reduce((obj, key) => {
    obj[key] = doc[key];
    return obj;
  }, {});;


saveAPI(doc, apiFileLocation);
