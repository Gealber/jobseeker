# Job seeker

The goal of this script is to perform scrapping looking for jobs. Doing it in a serious way. Websites that we will need to scrape:

1. [Google](http://google.com/).
2. [Ycombinator](https://news.ycombinator.com/jobs)
3. [Cloud Talent Solution API](https://cloud.google.com/talent-solution/job-search/v3/docs/best-practices)
4. [Linkedin Job Search](https://www.linkedin.com/jobs/search?keywords=Golang)

## Linkedin case

**Domain**
www.linkedin.com
**Path**
jobs/search
**Querys**
{
    "keyword": "Golang",
    "location": "United States",
    "position": 1,
    "pageNum": 0,
    "f_TPR": r604800,
    "f_WT": 2,
}

r604800 represents a week, 604800/60/60/24 = 7(days)
f_WT=2 represents the type of work, onsite = 1 and remote is 2

Examples: 

https://www.linkedin.com/jobs/search?keywords=Golang&location=United%20States&position=1&pageNum=0

https://www.linkedin.com/jobs/search?keywords=Golang&location=United%20States&locationId=&geoId=103644278&f_TPR=r604800&f_WT=2&position=1&pageNum=0

Clean the sample data scraped