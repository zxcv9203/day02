# day 02
## 쇼핑몰 오픈

우리는 방금 인터넷 쇼핑몰 창업을 시작했습니다.

먼저, 우리가 판매하는 상품과 상품의 가격을 지정하려고합니다.

상품의 이름과 가격을 적기위해 Product라는 구조체에 값을 저장하려고합니다.

Products 구조체는 다음과 같이 구성해야 합니다.
- `price int`

  물품의 가격을 지정합니다.

- `name string`

  해당 물품의 이름을 지정합니다.

Product 구조체는 외부에서 불러올 수 있지만 구조체 안의 price, name과 같은 필드는 직접적으로 접근할 수 없어야 합니다.

외부에서 price와 name에 접근하기 위한 메서드를 만들고 



Parcel의 필드에 직접적으로 접근해서는 안되며 외부 패키지에서 값을 접근할 수 있도록
택배 구조체의 값마다 값을 설정하는 메서드와 값을 가져오는 메서드를 작성해야합니다.

그리고 택배를 한번에 생성할 수 있도록 NewParcel 이라는 함수를 만들어봅시다.
`NewParcel(width, height, weight float64, content, etc string) (*Parcel, error)`

모든 함수와 메서드는 잘못된 값이 들어올 경우 적절한 에러처리가 되어있어야 합니다.

- 채점사항

1. Parcel 구조체가 외부로 export가 가능한지 확인해보세요. 다른패키지에서 접근이 불가능할 경우 FAIL입니다.

2. Parcel 구조체의 필드가 외부에서 직접 접근이 가능한지 확인해보세요. 예를들어, parcel.Width 같은 접근이 가능한지 확인해보세요
만약 접근이 가능하다면 FAIL 입니다.

3. 값을 가져오는 메서드와 설정하는 메서드가 잘 작성되었는지 확인하세요 메서드의 이름에 대해서 물어보세요. Go의 컨벤션에 맞게 작성
했는지 확인해보면 됩니다. 추가적으로 메서드에 에러처리가 잘 되어있는지 확인해보세요. 특정상황에서 문제를 일으킨다면 FAIL입니다. 예
를들어 width에 음수값을 넣어보고 에러를 반환하는지 확인해보세요
   
4. NewParcel이라는 함수가 잘 작성되어있는지 확인해보세요 직접 함수를 호출해보고 값이 잘 저장되고 잘못된 값 전달시 에러처리가 잘 되는지 확인해보세요
