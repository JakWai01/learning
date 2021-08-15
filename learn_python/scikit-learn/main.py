import joblib

# Target names for human-readable output
target_names = ['setosa', 'versicolor', 'virginica']

# Loading the pretrained model
classifier_knn = joblib.load('iris_classifier_knn.joblib')

# New sample to predict
custom_sample = [[5.9, 3.,  5.1, 1.8]]
preds = classifier_knn.predict(custom_sample)

# Print the target name(s) of the prediction(s)
pred_species = [target_names[p] for p in preds]
print("Custom predictions: ", pred_species)
