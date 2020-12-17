exports.handler = function(event, context, callback) {
    var res = {"event": event, "context": context}

    var headers = {};
    if (event.queryStringParameters["cache"].includes("true")) {
      headers = {"X-Nf-Experiment": "builder_func"};
    }


    callback(null, {
    statusCode: 200,
    body: JSON.stringify(res),
    headers: headers
    });
}
