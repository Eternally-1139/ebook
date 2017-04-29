import { KltshgPage } from './app.po';

describe('kltshg App', () => {
  let page: KltshgPage;

  beforeEach(() => {
    page = new KltshgPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
