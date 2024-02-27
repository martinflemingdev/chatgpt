import json
import os

def process_json_files(directory_path):
    # Loop through all files in the given directory
    for filename in os.listdir(directory_path):
        if filename.endswith(".json"):
            file_path = os.path.join(directory_path, filename)
            
            # Open and read the JSON file
            with open(file_path, 'r') as file:
                data = json.load(file)
            
            # Check if "Records" key exists
            if "Records" in data:
                for i, record in enumerate(data["Records"]):
                    if "body" in record:
                        # Deserialize the "body" content
                        body_content = json.loads(record["body"])
                        
                        # Define a new filename for the deserialized content
                        new_filename = f"{os.path.splitext(filename)[0]}_record_{i}.json"
                        new_file_path = os.path.join(directory_path, new_filename)
                        
                        # Save the deserialized "body" to a new file
                        with open(new_file_path, 'w') as new_file:
                            json.dump(body_content, new_file, indent=4)
            else:
                print(f'No "Records" found in {filename}')

def extract_message_from_files(directory_path):
    # Pattern to match files ending in "_record_0.json" or "_record_1.json"
    file_endings = ("_record_0.json", "_record_1.json")
    
    for filename in os.listdir(directory_path):
        if filename.endswith(file_endings):
            file_path = os.path.join(directory_path, filename)
            
            # Open and read the JSON file
            with open(file_path, 'r') as file:
                data = json.load(file)
            
            # Extract the "Message" field
            if "Message" in data:
                message_content = data["Message"]
                
                # Define a new filename for the message content
                new_filename = f"{os.path.splitext(filename)[0]}_message.json"
                new_file_path = os.path.join(directory_path, new_filename)
                
                # Save the "Message" to a new file, formatted for readability
                with open(new_file_path, 'w') as new_file:
                    # Check if message_content is a string that needs to be loaded as JSON
                    try:
                        message_json = json.loads(message_content)
                        json.dump(message_json, new_file, indent=4)
                    except json.JSONDecodeError:
                        # If message_content is not JSON, just write it as is
                        new_file.write(message_content)
            else:
                print(f'No "Message" found in {filename}')

if __name__ == "__main__":
    directory_path = '/path/to/your/directory'  # Update this to your directory path
    process_json_files(directory_path)
    extract_message_from_files(directory_path)
