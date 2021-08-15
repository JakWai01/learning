from graphene import ObjectType, String, Schema
from flask import Flask, request
import json

app = Flask(__name__)

class Query(ObjectType):
    hello = String(name=String(default_value="stranger"))
    goodbye = String()

    def resolve_hello(root, info, name):
        return f'Hello {name}!'

    def resolve_goodbye(root, info):
        return 'See ya!'

schema = Schema(query=Query)

@app.route('/graphql', methods=['POST'])
def graphql():
    data = json.loads(request.data)
    return json.dumps(schema.execute(data['query']).data)
