import requests
from concurrent.futures import ThreadPoolExecutor
from faker import Faker

fake = Faker()

# URL of the endpoint
url = "http://localhost:5000/usuarios"

# Function to create a user with random data
def create_user(_):
    nome = fake.name()
    email = fake.email()
    payload = {"nome": nome, "email": email}
    response = requests.post(url, json=payload)
    if response.status_code != 201:
        print(f"Created user. Status code: {response.status_code}, Response: {response.text}")

# Number of users to create
num_users = 50000

# Define the number of threads to use for concurrent requests
num_threads = min(10, 100)

# Create users concurrently using ThreadPoolExecutor
with ThreadPoolExecutor(max_workers=num_threads) as executor:
    # Use the same function for each user, but pass a dummy value (like underscore) as an argument
    # since the create_user function doesn't use the argument
    executor.map(create_user, range(num_users))
