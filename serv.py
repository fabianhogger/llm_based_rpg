from flask import Flask, request, jsonify
from ollama import chat
from ollama import ChatResponse

app = Flask(__name__)
 

@app.route('/generate', methods=['POST'])
def generate_text():
    try:
        data = request.json
        prompt = data.get("prompt", "")
        #max_length = data.get("max_length", 100)
        
        if not prompt:
            return jsonify({"error": "Prompt is required"}), 400

        response:ChatResponse = chat(model='llama3.2', messages=[
          {
            'role': 'user',
            'content': prompt,
          },
        ])
        print(response.message.content)
        return jsonify({"generated_text": response.message.content})
    except Exception as e:
        return jsonify({"error": str(e)}), 500

if __name__ == "__main__":
    app.run(debug=True)

