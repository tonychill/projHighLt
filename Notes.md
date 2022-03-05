# Post

Overall this reature request for a scheduler seems pretty straight forward from a design perspecting. 

## Conciderations
- It may be best to have the scheduler either run within a cluster or as a serverless component. The best option would be based on the service's demand. 
- In regards to ci/cd unit tests would be crucial (obviously). In addition an automated flow:   
  - code is written and build locally by the developer
  - the dev would pr against the main branch kicking off a git action or similar that would run integration and unit tests. 
  - An additional approach would be to add feature flags so that depoloyed features can be turnned off and on as needed. 
  - Once tests are good the feature branch can then be deployed either to the production environment in a canary fashion or to a staging environment for final qa testing. 
  
## Service Design

- Probably not the best approach for an excersise such as this one given the time constraints. However, I'm a huge fan of components and separating concerns. With that I took that approach where the routing, service and entity objects of the service wer separate. This allows for easy upgrades and verioning throughout the life of the service. 
- Utilities were used as a package however, I would lift this out to become a library on its own.
- From a broader perspective I would go for an event driven approach where the scheduler service would instead of calling the notification service or other services for that matter directly it would publish events. From there the notification service can subscribe to a broker and handle the envents that pertain to it. 
- Lastly, to handle the timers on creation the Reminder object would be created and persisted to a db with a start, end, durations and other options such as 30 min, 5 days, ect. Wether serverless or running in an ec2 instance the service will need to get a list of current running times as its spun up.  