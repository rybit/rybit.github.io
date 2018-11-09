
exports.handler = function(event, context, callback) {
    var res = "event: " + JSON.stringify(event)
    res += "\ncontext: " + JSON.stringify(context)
    res += "\n"

    callback(null, {
    statusCode: 200,
    body: res
    });
}
