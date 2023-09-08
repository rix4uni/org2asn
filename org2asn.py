import sys
import requests
from bs4 import BeautifulSoup
import json

def get_asn_info(org_name):
    response = requests.get(f"https://bgp.he.net/search?search%5Bsearch%5D={org_name}&commit=Search")
    soup = BeautifulSoup(response.text, 'html.parser')

    # Find the table containing the search results
    results_table = soup.find('div', id='search').find('table')

    # Extract the rows from the table, skipping the header
    rows = results_table.find_all('tr')[1:]

    results = []

    for row in rows:
        columns = row.find_all('td')
        # Extract the text from each column, using the .get_text() method for BeautifulSoup tags
        result = columns[0].get_text().strip()
        type_ = columns[1].get_text().strip()
        description = columns[2].get_text().strip()

        # Check if the type is not in ["Domain Name", "TLD"]
        if type_ not in ["Domain Name", "TLD"]:
            # Use a dictionary to store the extracted data for each row
            result_dict = {
                "Result": result,
                "Type": type_,
                "Description": description
            }
            results.append(result_dict)
    
    return results

def main():
    for org_name in sys.stdin:
        org_name = org_name.strip()
        results = get_asn_info(org_name)
        print(json.dumps(results, indent=2))

if __name__ == "__main__":
    main()
