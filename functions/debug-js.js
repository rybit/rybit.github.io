
exports.handler = function(event, context, callback) {
    var res = {"event": event, "context": context}


    callback(null, {
    statusCode: 200,
    body: JSON.stringify(res)
    });
}
