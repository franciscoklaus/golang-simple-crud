import requests
from concurrent.futures import ThreadPoolExecutor

# URL of the endpoint
base_url = "http://localhost:5000/usuarios"

# Function to delete a user by ID
def delete_user(user_id):
    url = f"{base_url}/{user_id}"
    response = requests.delete(url)
    print(f"Deleted user with ID {user_id}. Status code: {response.status_code}")

# Assuming you have a list of user IDs or obtain them from a previous response
user_ids = [x for x in range(2000)]

# Define the number of threads to use for concurrent requests
num_threads = min(10, 10)

# Delete users concurrently using ThreadPoolExecutor
with ThreadPoolExecutor(max_workers=num_threads) as executor:
    executor.map(delete_user, user_ids)