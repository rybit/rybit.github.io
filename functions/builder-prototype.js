exports.handler = function(event, context, callback) {
    var res = {"event": event, "context": context}

    var returnval = {
      statusCode: 200,
      body: JSON.stringify(res),
      headers: {"X-Nf-Experiment": "builder_func"}
    }
    console.log(returnval)

    callback(null, returnval);
}
