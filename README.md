# neogovapi
Neogov API Client written in Go

### About Cabarrus County
---
Cabarrus is an ever-growing county in the southcentral area of North Carolina. Cabarrus is part of the Charlotte/Concord/Gastonia NC-SC Metropolitan Statistical Area and has a population of about 210,000. Cabarrus is known for its rich stock car racing history and is home to Reed Gold Mine, the site of the first documented commercial gold find in the United States.

### About our team
---
The Business & Location Innovative Services (BLIS) team for Cabarrus County consists of six members:

+ Joseph Battinelli - Team Supervisor
+ Mark McIntyre - Software Developer
+ Landon Patterson - Software Developer
+ Aldair Balanzar - Software Developer
+ Jared Poe - GIS Analyst
+ Marci Jones - Software Developer

Our team is responsible for software development and support for the [County](https://www.cabarruscounty.us/departments/information-technology). We work under the direction of the Chief Information Officer.

### Auth
---
The Neogov API uses Basic Auth. This can be passed to the Client via the username and password parameters.

### Quick Examples
---

Spinning up a new client
```
ngClient := neogovapi.NewClient("yourusername", "yourpassword") //Store these securely
```

Querying Employees
```
	allEmps, err := ngClient.QueryEmployees(context.Background(), 10, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", allEmps)
```

Querying Departments
```
	allDepts, err := ngClient.QueryDepartments(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", allDepts)
```

Querying Evaluations
```
	allEvals, err := ngClient.QueryEvaulations(context.Background(), 10, 1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", allEvals)
```
### Coverage
---
#### Department

* GET /rest/departments
* GET /rest/departments/:departmentID

#### Employee

* GET /rest/employees
* GET /rest/employees/:employeeID
* GET /rest/employees/created
* GET /rest/employees/updated
* GET /rest/employees/:employeeID/evaluations
* GET /rest/employees/employeenumber/:employeeNumber


#### Evaluation

* GET /rest/evaluations
* GET /rest/evaluations/:evaluationsID

If an endpoint isn't covered, you can use the Send function in the client to handle these until we make them available. Rough example below

```
	emps := make([]map[string]interface{}, 0)

	r, err := http.NewRequestWithContext(context.Background(), "GET", "https://api.neogov.com/rest/employees?perpage=10&pagenumber=1", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ngClient.Send(r, &emps)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", emps)
```


More coverage to be added with each release.
