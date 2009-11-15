//
// Copyright 2009 Bill Casarin <billcasarin@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// generic.go - a collection of hacks that wouldn't be needed if there were
//              generics
package twitter

type tTwitterStatusDummy struct {
  Object tTwitterStatus;
}

type tTwitterTimelineDummy struct {
  Object []tTwitterStatus;
}

type tTwitterUserListDummy struct {
  Object []tTwitterUser;
}

type tTwitterUserDummy struct {
  Object User;
}

// I should probably put http_auth.go in here as well
