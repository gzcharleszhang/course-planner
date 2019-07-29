import requests
import json

"""
This script takes course information from the University of Waterloo Open Data API
and collects it into JSON files by subject in data/courses
"""

API_KEY = 'dbcdc71eb794474613d99a69db2bd3aa'
ALL_COURSES_URL = 'https://api.uwaterloo.ca/v2/courses.json'
COURSE_ID_URL = 'https://api.uwaterloo.ca/v2/courses/'

# Get all course IDs
courseIds = []
courses = requests.get(ALL_COURSES_URL, params={'key': API_KEY})
for course in courses.json()['data']:
    courseIds.append(course['course_id'])

# Get JSON for individual courses and group it with the appropriate subject
subjects = {}
for courseId in courseIds:
    url = COURSE_ID_URL + courseId + '.json'
    idRequest = requests.get(url, params={'key': API_KEY})

    subject = idRequest.json()['data']['subject']
    if subject in subjects:
        subjects[subject].append(idRequest.json()['data'])
    else:
        subjects[subject] = [idRequest.json()['data']]

# Write JSON for each subject into the appropriate file
for key in subjects.keys():
    jsonDict = {'data': subjects[key]}
    filePath = '../../../data/courses/' + key + '.json'
    with open(filePath, 'w') as file:
        json.dump(jsonDict, file)
