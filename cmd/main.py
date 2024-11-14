'''
Create local API to predict rating of a movies/film

input:
- NationID
- released_year
- num_of_episode
- time_per_episode
- DirectorID
- GernesID1
- GernesID2(optional)
- GernesID3(optional)

output:
- genres_list: {name of genres}
- predict_rating: [0:10]

Preprocessing:
Gernes:
1	Action
2	Adventure
3	Animation
4	Biography
5	Comedy
6	Crime
7	Drama
8	Fantasy
9	History
1	Horror
11	Music
12	Mystery
13	Romance
14	Sci-Fi
15	Sport
16	Thriller
17	War


if GernesID = (1,4,9)   ->  value: [1,0,0,1,0,0,0,0,1,0,0,0,0,0,0,0,0]
if GernesID = (4,6,14)  ->  value: [0,0,0,1,0,1,0,0,0,0,0,0,0,1,0,0,0]
'''
from flask import Flask, request, jsonify
from flask_cors import CORS
import numpy as np
import joblib

# Load the saved model and scaler (if scaling was used)
model = joblib.load("D:/Project/ReelPlay/cmd/movie_rating_predict_model.plk")

app = Flask(__name__)
CORS(app, resources={r"/predict": {"origins": "http://localhost:8080"}})

genres_list = ['Action','Adventure','Animation','Biography','Comedy','Crime','Drama','Fantasy','History','Horror','Music','Mystery','Romance','Sci-Fi','Sport','Thriller','War']
nation_weight_lvl1 = 1      # for nationID < 10
nation_weight_lvl2 = 10     # for nationID < 20
nation_weight_lvl3 = 100    # for nationID >=20

# Helper function to process input data
def process_input(nationId, directorId, releasedYear, numofeps, epLength, genre1, genre2, genre3):
    # Apply nation weight based on nationId
    if nationId < 10:
        nation_weighted = nationId * nation_weight_lvl1
        director_weighted = directorId * nation_weight_lvl1
    elif nationId < 20:
        nation_weighted = nationId * nation_weight_lvl2
        director_weighted = directorId * nation_weight_lvl2
    else:
        nation_weighted = nationId * nation_weight_lvl3
        director_weighted = directorId * nation_weight_lvl3

    # Create one-hot encoded genre vector (assuming 17 genres)
    genre_vector = [0] * 17
    for genre in [genre1, genre2, genre3]:
        if genre > 0:
            genre_vector[genre - 1] = 1

    # Combine all features into a single array
    features = [nation_weighted, releasedYear, numofeps, epLength, director_weighted] + genre_vector
    features = np.array(features).reshape(1, -1)  # Reshape for model input

    return features

# Route for predictions
@app.route('/predict', methods=['POST'])
def predict():
    # Parse input JSON
    data = request.get_json()
    print(data)
    # Extract input fields
    nationId = data.get("nationId")
    directorId = data.get("directorId")
    releasedYear = data.get("releasedYear")
    numofeps = data.get("numofeps")
    epLength = data.get("epLength")
    genre1 = data.get("genre1", 0)
    genre2 = data.get("genre2", 0)
    genre3 = data.get("genre3", 0)

    # Process input data for the model
    features = process_input(nationId, directorId, releasedYear, numofeps, epLength, genre1, genre2, genre3)
    # Run prediction
    prediction = model.predict(features)
    rating = prediction[0]  # Extract single prediction

    # Return prediction as JSON
    return jsonify({"predicted_rating": rating})

if __name__ == '__main__':
    app.run(debug=True)


