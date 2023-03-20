from typing import List


class DocumentAlreadyExistsError(Exception):
    pass


def get_athlete_by_doc_id(firestore_client, doc_id: str) -> dict:
    athlete_ref = firestore_client.collection("athletes").document(doc_id)
    athlete = athlete_ref.get()
    if athlete.exists:
        return athlete.to_dict()
    else:
        return {}

def create_athlete(
        firestore_client,
        doc_id: str,
        first_name: str,
        last_name: str,
        country: str) -> None:
    athlete_ref = firestore_client.collection("athletes").document(doc_id)
    athlete = athlete_ref.get()
    if athlete.exists:
        raise DocumentAlreadyExistsError()
    else:
        athlete_ref.set(
            {
                "first_name": first_name,
                "last_name": last_name,
                "country": country
            }
        )

def get_athletes_by_country(firestore_client, country: str) -> List[dict]:
    athletes_ref = firestore_client.collection("athletes")
    query = athletes_ref.where("country", "==", country)
    athletes_by_country = []
    for athlete in query.stream():
        athletes_by_country.append(athlete.to_dict())
    return athletes_by_country



def delete_activities_by_type(firestore_client, activity_type: str) -> None:
    activities_ref = firestore_client.collection("activities")
    activities = activities_ref.where("type", "==", activity_type).get()
    for activity in activities:
        key = activity.id
        activities_ref.document(key).delete()



def update_activity_map(firestore_client, doc_id: str) -> None:
    activities_ref = firestore_client.collection("activities")
    activity = activities_ref.document(doc_id)
    print(activity.get().to_dict())
    coordinates = activity.get().to_dict()["coordinates"]
    processed_coordinates = []
    for i in range(len(coordinates)):
        if i % 2 == 0:
            processed_coordinates.append(coordinates[i])
    activity.update({
        'processed_coordinates': processed_coordinates
    })


