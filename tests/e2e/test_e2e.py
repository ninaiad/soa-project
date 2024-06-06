import os
import httpx
import pytest


SERVER_ADDR = os.getenv("SERVER_ADDR", "localhost:8080")
USER_ENDPOINT = SERVER_ADDR + "/user"
POST_ENDPOINT = SERVER_ADDR + "/post/"

RETRYING = httpx.AsyncHTTPTransport(retries=5)


@pytest.mark.asyncio
async def test_user():
    async with httpx.AsyncClient(transport=RETRYING) as c:
        r = await c.post(USER_ENDPOINT + "/sign-in", json={"username": "nonexistant", "password": "testing"})
        assert r.status_code == 400

        r = await c.post(USER_ENDPOINT + "/sign-up", json={"username": "tester", "password": "testing"})
        assert r.status_code == 200
        r_json = r.json()
        assert "token" in r_json
        assert "user_id" in r_json
        user_id = r_json["user_id"]

        r = await c.post(USER_ENDPOINT + "/sign-in", json={"username": "tester", "password": "testing"})
        assert r.status_code == 200
        r_json = r.json()
        assert "token" in r_json
        assert "user_id" in r_json and r_json["user_id"] == user_id

        auth_header = {"Authorization": "Bearer " +
                       r_json["token"], "Content-Type": "application/json"}

        r = await c.put(USER_ENDPOINT + "/", json={"name": "Test!"}, headers=auth_header)
        assert r.status_code == 200
        r_json = r.json()
        assert "name" in r_json and r_json["name"] == "Test!"

        r = await c.put(USER_ENDPOINT + "/", json={"email": "test@test"}, headers=auth_header)
        assert r.status_code == 200
        r_json = r.json()
        assert "name" in r_json and r_json["name"] == "Test!"
        assert "email" in r_json and r_json["email"] == "test@test"


@pytest.mark.asyncio
async def test_posts():
    async with httpx.AsyncClient(transport=RETRYING) as c:
        r = await c.post(USER_ENDPOINT + "/sign-up", json={"username": "poster", "password": "testing"})
        assert r.status_code == 200
        r_json = r.json()
        assert "token" in r_json

        user_id = r_json["user_id"]
        auth_header = {"Authorization": "Bearer " +
                       r_json["token"], "Content-Type": "application/json"}

    async with httpx.AsyncClient(transport=RETRYING, headers=auth_header) as c:
        txt1 = "testing testing testing"
        r = await c.post(POST_ENDPOINT, json={"text": txt1})
        assert r.status_code == 200
        assert "post_id" in r.json()
        post_id0 = r.json()["post_id"]

        r = await c.get(POST_ENDPOINT, params={"id": post_id0, "author_id": user_id})
        assert r.status_code == 200
        assert "text" in r.json()
        assert "time_updated" in r.json()
        assert r.json()["text"] == txt1

        txt2 = "updated!"
        r = await c.put(POST_ENDPOINT, params={"id": post_id0}, json={"text": txt2})
        assert r.status_code == 200

        r = await c.get(POST_ENDPOINT, params={"id": post_id0, "author_id": user_id})
        assert r.status_code == 200
        assert "text" in r.json()
        assert r.json()["text"] == txt2

        r = await c.post(POST_ENDPOINT, json={"text": txt1})
        assert r.status_code == 200

        r = await c.get(USER_ENDPOINT + "/posts", params={"author_id": user_id, "page_num": 1, "page_size": 3})
        assert r.status_code == 200
        assert "posts" in r.json()
        posts = r.json()["posts"]
        assert len(posts) == 2
        assert (posts[0]["text"] == txt1 and posts[1]["text"] == txt2) \
            or (posts[1]["text"] == txt1 and posts[0]["text"] == txt2)
