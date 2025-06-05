import pickle
import json

with open("q_table.pkl", "rb") as f:
    data = pickle.load(f)

# แปลง key เป็น string แทนทั้งหมด
def key_to_str(key):
    return str(key)

converted_data = {key_to_str(k): v for k, v in data.items()}

# เขียนลงไฟล์ JSON
with open("q_table.json", "w") as f:
    json.dump(converted_data, f, indent=2)