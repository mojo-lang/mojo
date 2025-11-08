// Copyright 2022 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

@discriminator('@type')
@label_format('{}')
type Statement = AlterTableStmt @1
               | AnalyzeSmt @2
               | AttachStmt @3
               | BeginStmt @5
               | CommitStmt @6
               | CreateIndexStmt @10
               | CreateTableStmt @11
               | CreateTriggerStmt @12
               | CreateViewStmt @13
               | CreateVirtualTableStmt @14
               | DeleteStmt @17
               | DetachStmt @18
               | DropStmt @19
               | EndStmt  @30
               | InsertStmt @20
               | PragmaStmt @21
               | ReindexStmt @22
               | ReleaseStmt @23
               | RollbackStmt @24
               | SelectStmt @25
               | UpdateStmt @26
               | VacuumStmt @27

