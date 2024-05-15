package main

import "fmt"

const NMAX = 1000

const NMAX_Bank = 10

// AREA SUPER ADMIN VARIABEL
type Bank struct {
	name       string
	branch     [NMAX_Bank]Address
	nBranch    int
	uniqueCode int
	customers  [NMAX]Customer
	nCustomer  int
}

type WorldBank struct {
	Banks [NMAX_Bank]Bank
	nBank int
}

var worldBank WorldBank

type Credential struct {
	username, password string
}

var SuperAdmin = Credential{
	username: "1",
	password: "1",
}

// AREA NASABAH VARIABEL
type Customer struct {
	accountNumber int
	balance       int
	transactions  []Transaction
	nTransaction  int
	cardNumber    int
	PIN           int
	NIK           int
	name          string
	address       Address
	isSuspended   bool
}

type CustomerBank struct {
	customers [NMAX]Customer
	nCustomer int
}

var customerBank CustomerBank

type Transaction struct {
	transactionId          int
	senderBankCode         int
	senderAccountNumber    int
	recipientBankCode      int
	recipientAccountNumber int
	amount                 int
}

// AREA GENERAL VARIABEL
type Address struct {
	district string
	city     string
	province string
}

var startProgram bool = true

// AREA GENERAL FUNCTION

// main function to run the program
func main() {
	for startProgram {
		menu()
	}
}

// main menu function to show the menu
func menu() {
	var startMenu bool = true
	var guestChoice int
	fmt.Println("What user are u?")
	fmt.Println("1. Super Admin")
	fmt.Println("2. Bank Admin")
	fmt.Println("3. Customer")
	fmt.Println("4. Exit")

	for startMenu {
		fmt.Print("Input the option in number : ")
		fmt.Scan(&guestChoice)
		if guestChoice == 1 || guestChoice == 2 || guestChoice == 3 || guestChoice == 4 {
			startMenu = false
		} else {
			fmt.Println("Please input the right number")
			fmt.Print("\n")
			startMenu = true
		}
	}

	switch guestChoice {
	case 1:
		loginSuperAdmin()
		mainMenuSuperAdmin()
	case 2:
		// mainMenuAdminBank()
	case 3:
		mainMenuCustomer()
	case 4:
		exit()
	}
}

// exit function to exit the program
func exit() {
	var exit string
	fmt.Println("Are you sure want to exit? (Y/N)")

	fmt.Scan(&exit)
	for exit != "Y" && exit != "y" && exit != "N" && exit != "n" {
		fmt.Println("Please input the right option")
		fmt.Scan(&exit)
	}

	if exit == "Y" || exit == "y" {
		fmt.Println("Thank you for using our service")
		startProgram = false
	} else if exit == "N" || exit == "n" {
		fmt.Println("Thank you for using our service")
		fmt.Print("\n")
		menu()
	}
}

// AREA SUPER ADMIN FUNCTION

// mainMenuSuperAdmin function to show the main menu
var superAdminStart bool = true
var superAdminChoice int

func mainMenuSuperAdmin() {
	for superAdminStart {
		fmt.Println("\nSuper Admin Menu")
		fmt.Println("1. Insert Bank Data")
		fmt.Println("2. View Bank Data")
		fmt.Println("3. Edit Bank Data")
		fmt.Println("4. Delete Bank")
		fmt.Println("5. Logout")
		fmt.Print("Input : ")
		fmt.Scan(&superAdminChoice)
		for superAdminChoice >= 1 && superAdminChoice <= 5 {
			switch superAdminChoice {
			case 1:
				insertDataBank(&worldBank)
			case 2:
				viewDataBank(worldBank)
			case 3:
				editDataBank(&worldBank)
			case 4:
				deleteDataBank(&worldBank)
			case 5:
				logoutSuperAdmin()
			}
		}

	}
}

// logoutSuperAdmin function to logout the super admin
func logoutSuperAdmin() {
	var logout string
	fmt.Println("Are you sure want to logout? (Y/N)")

	fmt.Scan(&logout)
	for logout != "Y" && logout != "y" && logout != "N" && logout != "n" {
		fmt.Println("Please input the right option")
		fmt.Scan(&logout)
	}

	if logout == "Y" || logout == "y" {
		fmt.Println("Thank you for using our service")
		fmt.Print("\n")
		menu()
	} else if logout == "N" || logout == "n" {
		fmt.Print("\n")
		menu()
	}
}

// loginSuperAdmin function to login as super admin
func loginSuperAdmin() {
	var credential Credential
	var maxTrial int = 1
	fmt.Println("Welcome to Dashboard World Bank System")
	fmt.Println("--------------Login Menu--------------")
	fmt.Print("Please Input Username : ")
	fmt.Scan(&credential.username)
	fmt.Print("Please Input Password : ")
	fmt.Scan(&credential.password)
	// coba cek iz, ini mending or apa and, gua bikin sendiri, tapi ambigu sendiri wkwk
	for !isSuperAdmin(credential) && maxTrial < 4 {
		maxTrial++
		if maxTrial == 4 {
			fmt.Println("-- SORRY, YOU HAVE REACH THE MAXIMUM TRIAL FOR 3 TIMES --")
			fmt.Println("--                   Back To Menu                      --")
			fmt.Println("Loading.")
			fmt.Println("Loading..")
			fmt.Println("Loading...")
			fmt.Print("\n")
			maxTrial = 4
			menu()
		}

		fmt.Print("Login failed please try again\n")
		fmt.Print("Please Input Username : ")
		fmt.Scan(&credential.username)
		fmt.Print("Please Input Password : ")
		fmt.Scan(&credential.password)
	}

	fmt.Println("Login Success!")
}

// isSuperAdmin function to check the credential
func isSuperAdmin(credential Credential) bool {
	if credential.username == SuperAdmin.username && credential.password == SuperAdmin.password {
		return true
	}
	return false
}

// insertDataBank function to insert the bank data
func insertDataBank(worldBank *WorldBank) {
	var retry string = "y"
	var found bool

	var startInsertBank bool = true

	for startInsertBank {
		fmt.Println()
		fmt.Println("Bank", worldBank.nBank+1)

		for worldBank.Banks[worldBank.nBank].name == "" || len(worldBank.Banks[worldBank.nBank].name) < 3 {
			fmt.Print("Name : ")
			fmt.Scan(&worldBank.Banks[worldBank.nBank].name)
			if worldBank.Banks[worldBank.nBank].name == "" {
				fmt.Println("Name field may not be empty.")
			} else if len(worldBank.Banks[worldBank.nBank].name) < 3 {
				fmt.Println("Name should be contain more than 3 char.")
			}
		}

		var retryBranch string
		for retryBranch != "N" && retryBranch != "n" {
			fmt.Println()
			fmt.Println("Branch", worldBank.Banks[worldBank.nBank].nBranch+1)

			for worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district == "" || len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district) < 3 {
				fmt.Print("District : ")
				fmt.Scan(&worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district)
				if worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district == "" {
					fmt.Println("District field may not be empty.")
				} else if len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].district) < 3 {
					fmt.Println("District should be contain more than 3 char.")
				}
			}

			for worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city == "" || len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city) < 3 {
				fmt.Print("City : ")
				fmt.Scan(&worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city)
				if worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city == "" {
					fmt.Println("City field may not be empty.")
				} else if len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].city) < 3 {
					fmt.Println("City should be contain more than 3 char.")
				}
			}

			for worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province == "" || len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province) < 3 {
				fmt.Print("Province : ")
				fmt.Scan(&worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province)
				if worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province == "" {
					fmt.Println("Province field may not be empty.")
				} else if len(worldBank.Banks[worldBank.nBank].branch[worldBank.Banks[worldBank.nBank].nBranch].province) < 3 {
					fmt.Println("Province should be contain more than 3 char.")
				}
			}

			worldBank.Banks[worldBank.nBank].nBranch++
			fmt.Println("Branch Data Created Successfully")

			fmt.Println("\nWill You Create Branch Again? (Y/N)")
			fmt.Scan(&retryBranch)
			for retryBranch != "Y" && retryBranch != "N" && retryBranch != "y" && retryBranch != "n" {
				fmt.Scan(&retryBranch)
			}
		}

		for worldBank.Banks[worldBank.nBank].uniqueCode == 0 || worldBank.Banks[worldBank.nBank].uniqueCode < 100 || worldBank.Banks[worldBank.nBank].uniqueCode > 999 || found {
			found = false

			fmt.Print("Unique Code : ")
			fmt.Scan(&worldBank.Banks[worldBank.nBank].uniqueCode)

			for i := 0; i < worldBank.nBank && !found; i++ {
				if worldBank.Banks[i].uniqueCode == worldBank.Banks[worldBank.nBank].uniqueCode {
					found = true
				}
			}
			if worldBank.Banks[worldBank.nBank].uniqueCode == 0 {
				fmt.Println("Unique Code field may not be empty.")
			} else if worldBank.Banks[worldBank.nBank].uniqueCode < 100 || worldBank.Banks[worldBank.nBank].uniqueCode > 999 {
				fmt.Println("Unique Code should be greater than 99 and less than 1000")
			} else if found {
				fmt.Println("Unique Code is not available.")
			}
		}
		worldBank.Banks[worldBank.nBank].nCustomer = 0
		worldBank.nBank++

		fmt.Println("Bank Data Created Successfully")

		fmt.Println("\nWill Create Bank Data Again? (Y/N)")
		fmt.Scan(&retry)
		var startRetry bool = true
		for startRetry {
			if retry == "N" || retry == "n" {
				startRetry = false
				startInsertBank = false
				mainMenuSuperAdmin()
			} else if retry == "Y" || retry == "y" {
				startRetry = false
				startInsertBank = true
			} else {
				fmt.Println("Please input the right option")
				fmt.Scan(&retry)
			}
		}
	}
}

// viewDataBank function to view the bank data
func viewDataBank(worldBank WorldBank) {
	var startViewBank bool = true

	for startViewBank {
		if worldBank.nBank == 0 {
			fmt.Println("There is no data to be viewed.")
			startViewBank = false
			superAdminChoice = 0
		} else {
			fmt.Println()
			fmt.Println("Bank Data :")
			fmt.Printf("%-5s %-20s %-20s %-20s %-20s\n", "No", "Name", "Total Branch", "Unique Code", "Total Customer")
			for i := 0; i < worldBank.nBank; i++ {
				fmt.Printf("%-5d %-20s %-20d %-20d %-20d\n", i+1, worldBank.Banks[i].name, worldBank.Banks[i].nBranch, worldBank.Banks[i].uniqueCode, worldBank.Banks[i].nCustomer)
			}
			fmt.Println("Total Data : ", worldBank.nBank)
			startViewBank = false
			superAdminChoice = 0
		}
	}
}

// editDataBank function to edit the bank data
func editDataBank(worldBank *WorldBank) {
	var bankIndex, choice int
	var found bool

	if worldBank.nBank == 0 {
		fmt.Println("There is no data to be edited.")
		mainMenuSuperAdmin()
	}

	for bankIndex != -1 {
		viewDataBank(*worldBank)

		fmt.Print("Input Bank Number (input -1 for cancel) : ")
		fmt.Scan(&bankIndex)
		for (bankIndex-1 < 0 || bankIndex-1 > worldBank.nBank || !found) && bankIndex != -1 {
			for i := 0; i < worldBank.nBank && !found; i++ {
				if i == bankIndex-1 {
					found = true
				}
			}
			if bankIndex-1 < 0 || bankIndex-1 > worldBank.nBank {
				fmt.Println("Bank Number is not available.")
				fmt.Print("Input Bank Number (input -1 for cancel) : ")
				fmt.Scan(&bankIndex)
			}
		}

		for choice != 4 && bankIndex != -1 {
			fmt.Println("1. Edit Bank Name")
			fmt.Println("2. Edit Branch Data")
			fmt.Println("3. Edit Unique Code")
			fmt.Println("4. Save")
			fmt.Print("Input : ")
			fmt.Scan(&choice)
			for choice < 1 || choice > 4 {
				fmt.Print("Input : ")
				fmt.Scan(&choice)
			}

			switch choice {
			case 1:
				fmt.Println("Old Name :", worldBank.Banks[bankIndex-1].name)
				fmt.Print("New Name : ")
				fmt.Scan(&worldBank.Banks[bankIndex-1].name)
				for worldBank.Banks[bankIndex-1].name == "" || len(worldBank.Banks[bankIndex-1].name) < 3 {
					if worldBank.Banks[bankIndex-1].name == "" {
						fmt.Println("Name field may not be empty.")
					} else if len(worldBank.Banks[bankIndex-1].name) < 3 {
						fmt.Println("Name should be contain more than 3 char.")
					}
					fmt.Print("New Name : ")
					fmt.Scan(&worldBank.Banks[bankIndex-1].name)
				}
				fmt.Print("Bank Name Edited Successfully\n")
				viewDataBank(*worldBank)
				fmt.Println()

			case 2:
				var branchIndex, branchChoice int
				var foundBranch bool

				if worldBank.Banks[bankIndex-1].nBranch == 0 {
					fmt.Println("There is no branch data to be edited.")
				} else {
					for branchIndex != -1 {
						fmt.Println()
						fmt.Println("Branch Data :")
						fmt.Printf("%-5s %-20s %-20s %-20s\n", "No", "District", "City", "Province")
						for i := 0; i < worldBank.Banks[bankIndex-1].nBranch; i++ {
							fmt.Printf("%-5d %-20s %-20s %-20s\n", i+1, worldBank.Banks[bankIndex-1].branch[i].district, worldBank.Banks[bankIndex-1].branch[i].city, worldBank.Banks[bankIndex-1].branch[i].province)
						}
						fmt.Println("Total Data :", worldBank.Banks[bankIndex-1].nBranch)
						fmt.Print("Input Branch Number (input -1 for cancel) : ")
						fmt.Scan(&branchIndex)
						for (branchIndex-1 < 0 || branchIndex-1 > worldBank.Banks[bankIndex-1].nBranch || !foundBranch) && branchIndex != -1 {
							for i := 0; i < worldBank.Banks[bankIndex-1].nBranch && !foundBranch; i++ {
								if i == branchIndex-1 {
									foundBranch = true
								}
							}
							if branchIndex-1 < 0 || branchIndex-1 > worldBank.Banks[bankIndex-1].nBranch {
								fmt.Println("Branch Number is not available.")
								fmt.Print("Input Branch Number (input -1 for cancel) : ")

								fmt.Scan(&branchIndex)
							}
						}

						for branchChoice != 3 && branchIndex != -1 {
							fmt.Println("1. Edit District")
							fmt.Println("2. Edit City")
							fmt.Println("3. Edit Province")
							fmt.Println("4. Save")
							fmt.Print("Input : ")
							fmt.Scan(&branchChoice)
							for branchChoice < 1 || branchChoice > 4 {
								fmt.Print("Input : ")
								fmt.Scan(&branchChoice)
							}

							switch branchChoice {
							case 1:
								fmt.Println("Old District :", worldBank.Banks[bankIndex-1].branch[branchIndex-1].district)
								fmt.Print("New District : ")
								fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].district)
								for worldBank.Banks[bankIndex-1].branch[branchIndex-1].district == "" || len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].district) < 3 {
									if worldBank.Banks[bankIndex-1].branch[branchIndex-1].district == "" {
										fmt.Println("District field may not be empty.")
									} else if len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].district) < 3 {
										fmt.Println("District should be contain more than 3 char.")
									}
									fmt.Print("New District : ")
									fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].district)
								}
								fmt.Print("District Edited Successfully\n")
								viewDataBank(*worldBank)
								fmt.Println()

							case 2:
								fmt.Println("Old City :", worldBank.Banks[bankIndex-1].branch[branchIndex-1].city)
								fmt.Print("New City : ")
								fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].city)
								for worldBank.Banks[bankIndex-1].branch[branchIndex-1].city == "" || len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].city) < 3 {
									if worldBank.Banks[bankIndex-1].branch[branchIndex-1].city == "" {
										fmt.Println("City field may not be empty.")
									} else if len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].city) < 3 {
										fmt.Println("City should be contain more than 3 char.")
									}
									fmt.Print("New City : ")
									fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].city)
								}
								fmt.Print("City Edited Successfully\n")
								viewDataBank(*worldBank)
								fmt.Println()

							case 3:
								fmt.Println("Old Province :", worldBank.Banks[bankIndex-1].branch[branchIndex-1].province)
								fmt.Print("New Province : ")
								fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].province)
								for worldBank.Banks[bankIndex-1].branch[branchIndex-1].province == "" || len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].province) < 3 {
									if worldBank.Banks[bankIndex-1].branch[branchIndex-1].province == "" {
										fmt.Println("Province field may not be empty.")
									} else if len(worldBank.Banks[bankIndex-1].branch[branchIndex-1].province) < 3 {
										fmt.Println("Province should be contain more than 3 char.")
									}
									fmt.Print("New Province : ")
									fmt.Scan(&worldBank.Banks[bankIndex-1].branch[branchIndex-1].province)
								}
								fmt.Print("Province Edited Successfully\n")
								viewDataBank(*worldBank)
								fmt.Println()
							}
						}
					}
				}

			case 3:
				fmt.Println("Old Unique Code :", worldBank.Banks[bankIndex-1].uniqueCode)
				fmt.Print("New Unique Code : ")
				fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
				for worldBank.Banks[bankIndex-1].uniqueCode == 0 || worldBank.Banks[bankIndex-1].uniqueCode < 100 || worldBank.Banks[bankIndex-1].uniqueCode > 999 || found {
					found = false

					for i := 0; i < worldBank.nBank && !found; i++ {
						if worldBank.Banks[i].uniqueCode == worldBank.Banks[worldBank.nBank].uniqueCode {
							found = true
						}
					}

					if worldBank.Banks[bankIndex-1].uniqueCode == 0 {
						fmt.Println("Unique Code field may not be empty.")
						fmt.Print("New Unique Code : ")
						fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
					} else if worldBank.Banks[bankIndex-1].uniqueCode < 100 || worldBank.Banks[bankIndex-1].uniqueCode > 999 {
						fmt.Println("Unique Code should be greater than 99 and less than 1000")
						fmt.Print("New Unique Code : ")
						fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
					} else if found {
						fmt.Println("Unique Code is not available.")
						fmt.Print("New Unique Code : ")
						fmt.Scan(&worldBank.Banks[bankIndex-1].uniqueCode)
					}
				}
				fmt.Print("Unique Code Edited Successfully\n")
				viewDataBank(*worldBank)
				fmt.Println()
			}
		}
	}
}

// deleteDataBank function to delete the bank data
func deleteDataBank(worldBank *WorldBank) {
	var bankIndex int
	var found bool

	if worldBank.nBank == 0 {
		fmt.Println("There is no data to be deleted.")
		mainMenuSuperAdmin()
	}

	for bankIndex != -1 {
		viewDataBank(*worldBank)

		fmt.Print("Input Bank Number (input -1 for cancel) : ")
		fmt.Scan(&bankIndex)
		for (bankIndex-1 < 0 || bankIndex-1 > worldBank.nBank || !found) && bankIndex != -1 {
			for i := 0; i < worldBank.nBank && !found; i++ {
				if i == bankIndex-1 {
					found = true
				}
			}
			if bankIndex-1 < 0 || bankIndex-1 > worldBank.nBank {
				fmt.Println("Bank Number is not available.")
				fmt.Print("Input Bank Number (input -1 for cancel) : ")
				fmt.Scan(&bankIndex)
			}
		}

		if bankIndex != -1 {
			for i := bankIndex - 1; i < worldBank.nBank-1; i++ {
				worldBank.Banks[i] = worldBank.Banks[i+1]
			}
			worldBank.nBank--
			fmt.Println("Bank Data Deleted Successfully")
			mainMenuSuperAdmin()
			fmt.Println()
		}
	}
}

func initializeDummyCustomers() {
	dummyCustomers := []Customer{
		{accountNumber: 1001, balance: 0, PIN: 1234, NIK: 111111, name: "Dummy1", address: Address{district: "Dist1", city: "City1", province: "Prov1"}, isSuspended: false},
		{accountNumber: 1002, balance: 0, PIN: 1234, NIK: 111112, name: "Dummy2", address: Address{district: "Dist2", city: "City2", province: "Prov2"}, isSuspended: false},
		{accountNumber: 1003, balance: 0, PIN: 1234, NIK: 111113, name: "Dummy3", address: Address{district: "Dist3", city: "City3", province: "Prov3"}, isSuspended: false},
		{accountNumber: 1004, balance: 0, PIN: 1234, NIK: 111114, name: "Dummy4", address: Address{district: "Dist4", city: "City4", province: "Prov4"}, isSuspended: false},
		{accountNumber: 1005, balance: 0, PIN: 1234, NIK: 111115, name: "Dummy5", address: Address{district: "Dist5", city: "City5", province: "Prov5"}, isSuspended: false},
	}

	for i := 0; i < len(dummyCustomers); i++ {
		if customerBank.nCustomer < NMAX {
			customerBank.customers[customerBank.nCustomer] = dummyCustomers[i]
			customerBank.nCustomer++
		} else {
			fmt.Println("Customer bank is full, unable to add more customers.")
			break
		}
	}
}

func mainMenuCustomer() {
	initializeDummyCustomers()
	for {
		fmt.Println("Menu Customer:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Logout")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			login()
		case 2:
			register()
		case 3:
			fmt.Println("Logout berhasil.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func login() {
	var accountNumber, PIN int
	for attempts := 0; attempts < 3; attempts++ {
		fmt.Print("Masukkan nomor rekening: ")
		fmt.Scan(&accountNumber)
		fmt.Print("Masukkan PIN: ")
		fmt.Scan(&PIN)

		for i := 0; i < customerBank.nCustomer; i++ {
			if customerBank.customers[i].accountNumber == accountNumber && customerBank.customers[i].PIN == PIN {
				fmt.Println("Login berhasil.")
				customerMenu(&customerBank.customers[i])
				return
			}
		}

		fmt.Println("Nomor rekening atau PIN salah.")
	}

	fmt.Println("Akun Anda telah diblokir sementara karena terlalu banyak percobaan login.")
}

func register() {
	if customerBank.nCustomer >= NMAX {
		fmt.Println("Customer bank is full, unable to register new customer.")
		return
	}

	var customer Customer
	fmt.Print("Masukkan nomor rekening: ")
	fmt.Scan(&customer.accountNumber)
	fmt.Print("Masukkan saldo awal: ")
	fmt.Scan(&customer.balance)
	fmt.Print("Masukkan nomor kartu: ")
	fmt.Scan(&customer.cardNumber)
	fmt.Print("Masukkan PIN: ")
	fmt.Scan(&customer.PIN)
	fmt.Print("Masukkan NIK: ")
	fmt.Scan(&customer.NIK)
	fmt.Print("Masukkan nama: ")
	fmt.Scan(&customer.name)
	fmt.Print("Masukkan alamat (district, city, province): ")
	fmt.Scan(&customer.address.district, &customer.address.city, &customer.address.province)

	customerBank.customers[customerBank.nCustomer] = customer
	customerBank.nCustomer++

	fmt.Println("Registrasi berhasil.")
}

func customerMenu(customer *Customer) {
	for {
		fmt.Println("Menu Customer:")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Transfer")
		fmt.Println("3. Kembali")
		fmt.Print("Pilih menu: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewSaldo(customer)
		case 2:
			// Implementasi fungsi transfer di sini
			transfer(customer)
		case 3:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func viewSaldo(customer *Customer) {
	fmt.Printf("Saldo Anda adalah: %d\n", customer.balance)
}

func transfer(customer *Customer) {
	var recipientAccountNumber, amount int
	fmt.Print("Masukkan nomor rekening tujuan: ")
	fmt.Scan(&recipientAccountNumber)
	fmt.Print("Masukkan jumlah transfer: ")
	fmt.Scan(&amount)

	for i := 0; i < customerBank.nCustomer; i++ {
		if customerBank.customers[i].accountNumber == recipientAccountNumber {
			if customer.balance >= amount {
				customer.balance -= amount
				customerBank.customers[i].balance += amount
				fmt.Println("Transfer berhasil.")
				return
			} else {
				fmt.Println("Saldo tidak mencukupi.")
				return
			}
		}
	}

	fmt.Println("Nomor rekening tujuan tidak ditemukan.")
}
