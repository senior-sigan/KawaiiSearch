# -*- coding: utf-8 -*-
	
'''
Copy this as a config.py
'''

token = '-----------------------------------------------------------------------'
client_secret = '--------------------'
app_id = 1111111

def images_path(onwer_id):
	return 'data/images_{}'.format(onwer_id)

def info_path(onwer_id):
	return 'data/photos_{}.csv'.format(onwer_id)