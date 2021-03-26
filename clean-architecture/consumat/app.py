from repositories.database import Database
from use_cases.list_todos import ListTodos

data = Database()

result = str(ListTodos(data))

# split result to list
print(result)
