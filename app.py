import time

def transform():
    test_string = "hello"
    # Can sleep before returning response    
    #time.sleep(5)
    return test_string


from flask import Flask, request

app = Flask(__name__)

@app.route("/", methods=['POST'])
def transform_route():
    return transform()


if __name__ == "__main__":
    transform_route()
