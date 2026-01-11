#core/server.py
from core import app
from ariadne import gql, load_schema_from_path

schema = load_schema_from_path("./schema.graphql")


