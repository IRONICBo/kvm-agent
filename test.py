import argparse
import requests

def get_plugin_params(plugin_id, exec_result_id):
    url = "http://127.0.0.1:6000/api/plug/param?execResultId=" + str(exec_result_id) + "&plugId=" + str(plugin_id)
    # params = {"execResultId": exec_result_id, "plugId": plugin_id}
    print("Fetching plugin params from: " + url)
    response = requests.get(url)
    if response.status_code == 200:
        print("Successfully fetched plugin params:")
        print(response.json())
    else:
        print("Failed to fetch plugin params")

def send_plugin_result(plugin_id, exec_result_id, result_text, result_file_path=""):
    url = "http://127.0.0.1:6000/api/plug/result"
    data = {
        "execResultId": exec_result_id,
        "plugId": plugin_id,
        "plugResult": result_text,
        "plugResultFilePath": result_file_path
    }
    response = requests.post(url, json=data)
    if response.status_code == 200:
        print("Successfully sent plugin result")
    else:
        print("Failed to send plugin result")

def main():
    parser = argparse.ArgumentParser(description='Client for plugin parameter retrieval and result sending.')
    parser.add_argument('--pluginId', type=int, required=True, help='Plugin ID')
    parser.add_argument('--execResultId', type=int, required=True, help='Execution Result ID')
    args = parser.parse_args()

    # Fetch plugin parameters
    get_plugin_params(args.pluginId, args.execResultId)

    # Example: Send a plugin result. Replace "Your result text here" with actual result text.
    send_plugin_result(args.pluginId, args.execResultId, "Your result text here")

if __name__ == "__main__":
    main()
