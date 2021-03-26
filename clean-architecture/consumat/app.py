from repositories.database import Database
from use_cases.list_todos import ListTodos

data = Database()

print(ListTodos(data))
