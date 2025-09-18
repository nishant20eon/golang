# Functions across Packages

## 🔹 10 Real-World Practice Questions (Functions across Packages)

### 1. **Math Utils – Square & Cube**

- Package `mathutils` has a function `Square(n int) int`.
- In `main`, write a function `Cube(n int) int` that calls `mathutils.Square` internally to calculate cube.

---

### 2. **Greeting Service**

- Package `greet` has function `Hello(name string) string`.
- In `main`, create function `WelcomeUser(name string)` that calls `Hello` and appends `"Welcome to our system"`.

---

### 3. **Temperature Converter**

- Package `converter` has `CelsiusToFahrenheit(c float64) float64`.
- In `main`, create function `WeatherReport(celsius float64) string` that calls the converter and returns a formatted string like `"Today’s temp is XX°F"`.

---

### 4. **Payment Processing (Multi Param, Single Return)**

- Package `payment` has function `Discount(price float64, percent float64) float64`.
- In `main`, write function `FinalAmount(price float64, discount float64) float64` that calls `Discount` to compute net price.

---

### 5. **User Login (Multi Param, Multi Return)**

- Package `auth` has `ValidateUser(username, password string) (bool, error)`.
- In `main`, create function `Login(username, password string) string` that calls `ValidateUser` and returns `"Login Success"` or `"Login Failed"`.

---

### 6. **Bank Account (Value Receiver)**

- Package `bank` has struct `Account{Balance float64}` with method `(a Account) ShowBalance() float64`.
- In `main`, create function `CheckBalance(a bank.Account)` that calls `ShowBalance` and prints result.

---

### 7. **Bank Account (Pointer Receiver)**

- Package `bank` also has method `(a *Account) Deposit(amount float64)`.
- In `main`, create function `SalaryCredit(a *bank.Account, salary float64)` that calls `Deposit`.

---

### 8. **Student Grading**

- Package `student` has function `Average(marks []int) float64`.
- In `main`, write function `ResultStatus(marks []int) string` that calls `Average` and returns `"Pass"` if avg ≥ 40, else `"Fail"`.

---

### 9. **E-commerce Order (Multi Return)**

- Package `order` has function `CalculateTotal(prices []float64) (float64, int)` → returns total and item count.
- In `main`, write function `OrderSummary(prices []float64) string` that calls `CalculateTotal` and prints formatted summary.

---

### 10. **Car Service (Pointer Receiver + Package Function Call)**

- Package `car` has struct `Car{Brand string, Mileage int}` with method `(c *Car) AddMileage(miles int)`.
- In `main`, create function `Travel(c *car.Car, miles int)` that calls `AddMileage` and then calls another package function `report.MileageReport(c)` to print updated info.

---

✅ These 10 cover **all combinations** you asked:

- single param, multi param
- single return, double return
- value receiver, pointer receiver
- cross-package calls

[Solution - 1](Solution%20-%201%20272d9598f0ab8003a63fffae0c87a2a2.md)

[Solution - 2](Solution%20-%202%20272d9598f0ab80e9b4dadbf19046ca9b.md)