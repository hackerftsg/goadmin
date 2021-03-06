package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
	"log"
	"net/http"
	"io/ioutil"
	"regexp"
	"strconv"
)

const (
	appauthor = "d3z3n0v3"
	appversion = "1.0.0"
	apptitle = "GoAdmin Scanner"
	TRUE = 1
	FALSE = 0
)

type Path struct {
	path []string
}

func (admin * Path) new() {
	admin.path = []string{"admin/","admin/login.php","Admin/","Admin/Login.php","Administrador/","wp/wp-login.php","wp-login.php","administrator/","moderator/","webadmin/","adminarea/","bb-admin/","adminLogin/","admin_area/","panel-administracion/","instadmin/",
		"memberadmin/","administratorlogin/","adm/","account.cfm","admin/account.cfm","admin/index.cfm","admin/login.cfm","admin/admin.cfm",
		"admin_area/admin.cfm","admin_area/login.cfm","admin/account.html","admin/index.html","admin/login.html","admin/admin.html",
		"admin_area/admin.html","admin_area/login.html","admin_area/index.html","admin_area/index.cfm","bb-admin/index.cfm","bb-admin/login.cfm","bb-admin/admin.cfm",
		"bb-admin/index.html","bb-admin/login.html","bb-admin/admin.html","admin/home.html","admin/controlpanel.html","admin.html","admin/cp.html","cp.html",
		"administrator/index.html","administrator/login.html","administrator/account.html","administrator.html","login.html","modelsearch/login.html","moderator.html",
		"moderator/login.html","moderator/admin.html","account.html","controlpanel.html","admincontrol.html","admin_login.html","panel-administracion/login.html",
		"admin/home.cfm","admin/controlpanel.cfm","admin.cfm","pages/admin/admin-login.cfm","admin/admin-login.cfm","admin-login.cfm","admin/cp.cfm","cp.cfm",
		"administrator/account.cfm","administrator.cfm","login.cfm","modelsearch/login.cfm","moderator.cfm","moderator/login.cfm","administrator/login.cfm",
		"moderator/admin.cfm","controlpanel.cfm","admin/account.html","adminpanel.html","webadmin.html","pages/admin/admin-login.html","admin/admin-login.html",
		"webadmin/index.html","webadmin/admin.html","webadmin/login.html","user.cfm","user.html","admincp/index.cfm","admincp/login.cfm","admincp/index.html",
		"admin/adminLogin.html","adminLogin.html","admin/adminLogin.html","home.html","adminarea/index.html","adminarea/admin.html","adminarea/login.html",
		"panel-administracion/index.html","panel-administracion/admin.html","modelsearch/index.html","modelsearch/admin.html","admin/admin_login.html",
		"admincontrol/login.html","adm/index.html","adm.html","admincontrol.cfm","admin/account.cfm","adminpanel.cfm","webadmin.cfm","webadmin/index.cfm",
		"webadmin/admin.cfm","webadmin/login.cfm","admin/admin_login.cfm","admin_login.cfm","panel-administracion/login.cfm","adminLogin.cfm",
		"admin/adminLogin.cfm","home.cfm","admin.cfm","adminarea/index.cfm","adminarea/admin.cfm","adminarea/login.cfm","admin-login.html",
		"panel-administracion/index.cfm","panel-administracion/admin.cfm","modelsearch/index.cfm","modelsearch/admin.cfm","administrator/index.cfm",
		"admincontrol/login.cfm","adm/admloginuser.cfm","admloginuser.cfm","admin2.cfm","admin2/login.cfm","admin2/index.cfm","adm/index.cfm",
		"adm.cfm","affiliate.cfm","adm_auth.cfm","memberadmin.cfm","administratorlogin.cfm","siteadmin/login.cfm","siteadmin/index.cfm","siteadmin/login.html"}
}

type Scan struct {
	url, path string
}

func (scan * Scan) new(arg string) {
	scan.url = arg
}

func (scan * Scan) setpath(arg string) {
	scan.path = arg
}

func (scan * Scan) status() []string {
	url := fmt.Sprintf("%s%s", scan.url, scan.path)
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	strbody, strcode := string(body), strconv.Itoa(resp.StatusCode)

	return []string{strbody, strcode}
}

func (scan * Scan) response() int {
	result := scan.status()
	body, code := result[0], result[1]

	intcode, _ := strconv.Atoi(code)

	x, y := 1, 1

	if (intcode != 200) && (intcode != 201) && (intcode != 202) && (intcode != 203) {
		x--
	}

	matched, _ := regexp.MatchString("<(FORM|form).+(POST|post).+>", body)

	if matched == false {
		y--
	}

	z := x + y

	if z > TRUE {
		return TRUE
	}
	return FALSE
}

func isurl(arg string) bool {
	if (strings.Contains(arg, "http://") || strings.Contains(arg, "https://")) && arg[len(arg)-1:] == "/" {
		return true
	}
	return false
}

func main() {

	if len(os.Args) > 1 {
		path := Path{}
		path.new()

		title := fmt.Sprintf("%s created by %s version %s\n", apptitle, appauthor, appversion)
		fmt.Println(title);

		scan := Scan{}

		arg := os.Args[1]

		if isurl(arg) {
			scan.new(arg)
			for _, element := range path.path {
				scan.setpath(element)
				if scan.response() == TRUE {
					response := fmt.Sprintf("[+] Found! %s%s", arg, element)
					fmt.Println(response)
					break
				}
			}
		} else if _, err := os.Stat(arg); !os.IsNotExist(err) {
			file, err := os.Open(arg)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				if isurl(scanner.Text()) {
					scan.new(scanner.Text())
					for _, element := range path.path {
						scan.setpath(element)
						if scan.response() == TRUE {
							response := fmt.Sprintf("[+] Found! %s%s", scanner.Text(), element)
							fmt.Println(response)
							break
						}
					}
				}
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		} else {
			err := fmt.Sprintf("A argumento ('%s') requisitado é inválida.\n\nURL Deve conter http:// ou https:// e / no final da url.\n\nOu o arquivo especificado não existe.", arg)
			fmt.Println(err)
		}
	} else {
		fmt.Println("Usage: go run goadmin.go http://target.com/ OR list.txt")
	}
}
