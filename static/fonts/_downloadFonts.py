import requests
from tqdm import tqdm

def DownloadFont(url: str) -> bool:
	filename = url.split('/')[-1]
	r = requests.get(url)
	if r.status_code == 200:
		with open(filename, 'wb') as file:
			file.write(r.content)

	return 

def Main() -> None:
	urls = []
	with open('_urls.txt', 'r') as file:
		urls = [line.strip() for line in file]

	for url in tqdm(urls, total=len(urls)):
		DownloadFont(url=url)


if __name__ == '__main__':
	Main()