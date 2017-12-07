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

ssh_site:
	gcloud compute --project "mltest-180907" ssh --zone "europe-west1-d" "kawaii-search"

copy_site:
	gcloud compute ssh --command "rm -rf ~/images_similarity/src" --project "mltest-180907" --zone "europe-west1-d" "kawaii-search"
	gcloud compute ssh --command "rm -rf ~/images_similarity/templates" --project "mltest-180907" --zone "europe-west1-d" "kawaii-search"
	gcloud compute scp templates "kawaii-search":~/images_similarity/templates --recurse --zone europe-west1-d --compress
	gcloud compute scp src "kawaii-search":~/images_similarity/src --recurse --zone europe-west1-d --compress

copy_site_data:
	gcloud compute scp submission/images_order_-70232735.csv "kawaii-search":~/images_similarity/submission/images_order_-70232735.csv --zone europe-west1-d
	gcloud compute scp data/photos_-70232735.csv "kawaii-search":~/images_similarity/data/photos_-70232735.csv --zone europe-west1-d
	gcloud compute scp submission/images_vec_-70232735.npz "kawaii-search":~/images_similarity/submission/images_vec_-70232735.npz --zone europe-west1-d