택배 배송업체

만든 택배를 배송하기위해 택배 배송 업체를 만듭시다.

택배 배송 업체의 정보를 담은 구조체 Cj를 만듭니다.

Cj 구조체는 택배 구조체를 임베딩하고 있어야 합니다.

Cj 구조체는 다음과 같은 내용을 포함해야 합니다.

- 송신인을 저장할 `to` 변수

- 수신인을 저장할 `from` 변수

- 택배 정보를 보는 `Info` 메서드

- 택배를 발송하는 메시지를 출력하는 `Send` 메서드

Info 메서드는 임베딩한 택배의 내용을 출력해주며 만약 내용이 비어있을 경우 해당 필드에 값을 입력하라는 메시지가 나옵니다.

Send 메서드는 보내는 사람과 받는 사람이 존재하는지 확인하고 존재하면 Send에 성공했다는 메시지를 출력하고 없다면 받는 사람과 보내는 사람을 적은 후 보내라고 명시합니다.

평가항목

Cj구조체가 Parcel 구조체를 임베딩하고 있는지 확인해야 합니다. 그리고 Parcel 구조체에 만들었던 메서드가 잘 작동하는지 확인해야 합니다. 작동하지 않는다면 FAIL입니다.

택배를 보내는사람과 받는 사람이 필드로 선언되어 있는지 확인합니다. 해당 필드들은 외부로 노출되어있어서는 안됩니다.

