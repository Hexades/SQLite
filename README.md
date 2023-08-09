# SQLit Hexade
SQLite Repository 
The Hexades architecture uses a local bus with events which contain (a) a channel to send data back to the caller (errors and data) and (b) a higher order function to execute. The function is executed receving the repository and the data. This function is customizable to a client's specific requirements and would be the equivalent, perhaps, of an adapter in pure imperative hexagonal architecture which doesn't use messaging. 
```
type Insert struct {
    bus.Event
    value      any
    insertFunc InsertFunction
}
```
This is the extent of the repository logic! It receives the event on a channel, verifies it is an SQLite event, and then simply calls execute on the event passing itself in. The Execute method in turn calls the InsertFunc and executes the logic. 

```
func (r *repository) OnRepositoryEvent(repositoryChannel <-chan bus.RepositoryEvent) {

    for repoEvent := range repositoryChannel {
        switch evt := repoEvent.(type) {
        case SQLiteEvent:
            evt.Execute(r)
        }
    }
}
```

The basic insert  function. This where things can get interesting in the future. The event, the repository and all the plumbing stay the same but if I want some different handling of the insert. An update is even a better example, perhaps, as I may need to look up some values first before updating. In any case, this is  the highly customizable part. Just  inject a different function into the event  and you change the behavior. 
```
var BasicInsertFunc = func(event *Insert, repo *repository) {
    value := event.value
    //TODO Remove this after initial development
    repo.db.AutoMigrate(&value)
    tx := repo.db.Create(value)
    event.Send(bus.NewResponse(value, tx.Error))
}
```

The final step in the function is to call Send on the event. That actually puts the response on a channel in the event and the receiver can listen for it and block while all the other processing happens concurrently or in parallel.
