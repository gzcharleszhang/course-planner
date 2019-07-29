import requests
import json

"""
This script takes course information from the University of Waterloo Open Data API
and collects it into JSON files by subject in data/courses
"""

API_KEY = 'dbcdc71eb794474613d99a69db2bd3aa'
ALL_COURSES_URL = 'https://api.uwaterloo.ca/v2/courses.json'
COURSE_ID_URL = 'https://api.uwaterloo.ca/v2/courses/'
JSON_PATH = '../data/courses/'

# Get all course IDs
course_ids = []
courses = requests.get(ALL_COURSES_URL, params={'key': API_KEY})
for course in courses.json()['data']:
    course_ids.append(course['course_id'])

# Get JSON for individual courses and group it with the appropriate subject
subjects = {}
for course_id in course_ids:
    url = COURSE_ID_URL + course_id + '.json'
    id_request = requests.get(url, params={'key': API_KEY})

    subject = id_request.json()['data']['subject']
    if subject in subjects:
        subjects[subject].append(id_request.json()['data'])
    else:
        subjects[subject] = [id_request.json()['data']]

# Write JSON for each subject into the appropriate file
for key in subjects.keys():
    json_dict = {'data': subjects[key]}
    file_path = JSON_PATH + key + '.json'
    with open(file_path, 'w') as file:
        json.dump(json_dict, file)
