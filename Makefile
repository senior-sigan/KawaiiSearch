copy:
	gcloud compute ssh --command "rm -rf ~/images_similarity/src" --project "mltest-180907" --zone "europe-west1-d" "keras--tf-gpu"
	gcloud compute scp src "keras--tf-gpu":~/images_similarity/src --recurse --zone europe-west1-d --compress

copy_data:
	gcloud compute scp data/images "keras--tf-gpu":~/images_similarity/data/images --recurse --zone europe-west1-d

ssh:
	gcloud compute ssh --project "mltest-180907" --zone "europe-west1-d" "keras--tf-gpu"

shutdown:
	gcloud compute ssh --command "sudo shutdown -h now" --project "mltest-180907" --zone "europe-west1-d" "keras--tf-gpu"

download:
	gcloud compute scp "keras--tf-gpu":~/images_similarity/submission submission/ --recurse --compress --zone europe-west1-d

start:
	gcloud compute instances start keras--tf-gpu --zone=europe-west1-d