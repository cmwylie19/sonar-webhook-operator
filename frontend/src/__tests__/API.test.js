import API from "../utils";
import moxios from "moxios";

export const mockHealthCheck = {
  data: { "x-forwarded-for": "" },
};

describe("API", () => {
  const axiosInstance = new API();
  beforeEach(() => {
    moxios.install(axiosInstance.instance);
  });

  afterEach(() => {
    moxios.uninstall(axiosInstance.instance);
  });

  it("should GET the healthcheck endpoint", async () => {
    moxios.wait(() => {
      const request = moxios.requests.mostRecent();
      request.respondWith({ status: 200, response: mockHealthCheck });
    });
    const response = await axiosInstance.pingHealthCheck();
    console.log(response.data);
    expect(response.data).toEqual(mockHealthCheck);
  });
});
