# [ソフトウェアコンテスト 2024](https://www.ogis-ri.co.jp/otc/contest/)

これはソフトウェアコンテスト 2024 のためのリポジトリです。
フロントエンドは、go で実装し、通信方式は gRPC を採用しています

# 技術選定

| Technology                     | Libraries and tools used |
| ------------------------------ | ------------------------ |
| Language                       | go1.22.45                |
| Local environment construction | Docker                   |
| ORM                            |
| Migration                      | Goose                    |
| Seeds                          | Self-implementation      |
| Schema                         | Graph QL                 |

## インデックス

- 🌳[環境構築](./docs/setUp.md)
- 🍏[DB 更新](./docs/migration.md)
- 📗[ディレクトリー構成](./docs/strucure.md)
- 🍓[api 作成手順](./docs/proto.md)
