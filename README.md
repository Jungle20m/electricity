# Go Clean Architecture Template

Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.

[![Go Report Card](https://goreportcard.com/badge/github.com/evrone/go-clean-template)](https://goreportcard.com/report/github.com/evrone/go-clean-template)
[![License](https://img.shields.io/github/license/evrone/go-clean-template.svg)](https://github.com/evrone/go-clean-template/blob/master/LICENSE)
[![Release](https://img.shields.io/github/v/release/evrone/go-clean-template.svg)](https://github.com/evrone/go-clean-template/releases/)
[![codecov](https://codecov.io/gh/evrone/go-clean-template/branch/master/graph/badge.svg?token=XE3E0X3EVQ)](https://codecov.io/gh/evrone/go-clean-template)


## Contents
- [Clean Architecture](#clean-architecture)
- [Project Structure](#project-structure)
- [SDK](#sdk)
  - [x] Mysql
  - [x] Logger
  - [x] Rest Api
  - [x] GRPC


Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.



## Clean Architecture
***Clean Architecture*** là một kiến trúc ứng dụng rất nổi tiếng dựa trên nguyên lý loại bỏ sự lệ thuộc giữa các đối tượng
cũng như các layer trong ứng dụng. Nguyên lý này kế thừa và phát triển dựa trên **Dependency Inversion** - nguyên lý nổi
tiếng trong SOLID.

Trong kiến trúc Clean Architecture bao gồm 4 layer được đại diện thông qua các vòng tròn đồng tâm. Các vòng tròn ở trong 
sẽ không hề biết gì về các vòng tròn bên ngoài. Nguyên tắc "hướng tâm" này được minh hoạ như sau:

![Clearn Architecture 1](docs/img/ca1.png)

### Entities
***Entities*** là layer trong cùng, cũng là layer quan trọng nhất. Entity chính là các thực thể hay từng đối tượng cụ thể và 
các rule business logic của nó. Trong **OOP**, đây chính là **Object** cùng với các **method** và **properties** tuân thủ nguyên tắc 
**Encapsulation** - chỉ bên trong Object mới có thể thay đổi trạng thái (**State**) của chính nó.

VD: Trong object **Person** thì thuộc tính **age** không thể bé hơn 1. Nếu cần thay đổi age, chúng ta phải viết hàm **public setAge**, 
hàm này cũng chịu trách nhiệm check điều kiện liên quan tới **age**.

Các business logic của layer Entities sẽ không quan tâm hay lệ thuộc vào các business logic ở các layer bên ngoài như Use Cases. 
Giả sử với trường hợp người dùng phải từ 18 tuổi trở lên mới được phép tạo tài khoản thì rule thuộc tính Age trong Entities vẫn không đổi.

### Use Cases

### Interface Adapters

### Frameworks & Drivers


![Clearn Architecture 2](docs/img/ca2.png)

Where can I get some?
There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which don't look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isn't anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc.

## Project Structure
    .
    ├── .local                    
    ├── common                     
    ├── config                     
    │   ├── config.go             
    │   ├── config.yaml            
    │   └── config.sample.yaml           
    ├── docs                       
    ├── internal
    │   ├── http-server      
    │   ├── grpc-server      
    │   └── modules 
    │       ├── order
    │       │   ├── business
    │       │   ├── transport
    │       │   ├── model
    │       │   └── modules
    │       └── payment
    ├── sdk                      
    │   ├── mysql                  
    │   ├── logger                 
    │   ├── redis
    │   └── httpserver  
    ├── main.go                 
    └── README.md

### common
- Khai báo các hằng số
- Khởi tạo AppContext sử dụng toàn bộ chương trình

### config
- `config.go` đọc config từ file `config.yaml`
- `config.yaml`
- `config.sample.yaml`

### docs
Chứa các file swagger được sinh tự động bằng thư viện [swag](https://github.com/swaggo/swag).

### internal
Chứa source code chính của toàn bộ project

### sdk


## SDK

Section 1.10.32 of "de Finibus Bonorum et Malorum", written by Cicero in 45 BC
"Sed ut perspiciatis unde omnis iste natus error sit voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur?"