# Changelog

All notable changes to this project will be documented in this file. See [standard-version](https://github.com/conventional-changelog/standard-version) for commit guidelines.

### [21.2.1](https://github.com/mcauto/todolist-api/compare/v21.2.0...v21.2.1) (2022-07-27)


### Bug Fixes

* **makefile:** go workspace 사용 시 하위 패키지를 ./...으로 찾지 못하는 버그 수정 ([62fcc32](https://github.com/mcauto/todolist-api/commit/62fcc323dbc6058896b038a4f817f040e32eb556))

## [21.2.0](https://github.com/mcauto/todolist-api/compare/v21.1.1...v21.2.0) (2022-03-17)


### Features

* apply multi-module (go1.18) ([c86b835](https://github.com/mcauto/todolist-api/commit/c86b835ac945277bd1bd6eacda767c31e52bd6c5))

### [21.1.1](https://github.com/mcauto/todolist-api/compare/v21.1.0...v21.1.1) (2021-12-05)


### Bug Fixes

* **todo:** insert, fetch, update 결과 처리 방식 변경 ([b29a92a](https://github.com/mcauto/todolist-api/commit/b29a92abf33262459cee7c20416fa5a2d3230a01))

## 21.1.0 (2021-12-04)


### Features

* **changelog:** 버전 자동관리 추가 ([9aea966](https://github.com/mcauto/todolist-api/commit/9aea966207e5b341a946b71d633fbe72b47f8bec))
* config, cmd, docs 추가 ([b510101](https://github.com/mcauto/todolist-api/commit/b5101015c4838d6f9bcafc365eaf061cae4bc3f1))
* **database:** database 연동 ([f2732e2](https://github.com/mcauto/todolist-api/commit/f2732e287f50406999d14c05b0968ee4675b2afd))
* **database:** 개발환경 시 sql logger 추가 ([9e92386](https://github.com/mcauto/todolist-api/commit/9e92386e7f0fc37f1ca9a25c6d408d0a6a1a3505))
* **delivery:** add web server, move cmd ([a0922d7](https://github.com/mcauto/todolist-api/commit/a0922d76b2dfff47a7e0784ac39d5be9a1286487))
* **docker:** 컨테이너 빌드 추가 ([4d3e5a5](https://github.com/mcauto/todolist-api/commit/4d3e5a5ae36487bc67245a65f1f46e3d981f3a81))
* **repository:** add mysql repository (gorm) ([c43ac38](https://github.com/mcauto/todolist-api/commit/c43ac389a4dc5efaf6f991e16ca9885d8bd8ec98))
* **repository:** debug 모드인 경우 sqlite로 실행 ([0937f65](https://github.com/mcauto/todolist-api/commit/0937f651e9223d3b8f19254749090626493c3c0e))
* **web:** todo crud 구현 ([bd7d4bc](https://github.com/mcauto/todolist-api/commit/bd7d4bc641093c7722220f36516c2b37194436e9))


### Bug Fixes

* **repository:** mysql -> dbms ([61e2a70](https://github.com/mcauto/todolist-api/commit/61e2a7089d91c58a2baa36a88d5720b613d5a194))
