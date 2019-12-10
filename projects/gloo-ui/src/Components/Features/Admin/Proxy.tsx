import styled from '@emotion/styled';
import { ReactComponent as ProxyLogo } from 'assets/proxy-icon.svg';
import { ConfigDisplayer } from 'Components/Common/DisplayOnly/ConfigDisplayer';
import { FileDownloadLink } from 'Components/Common/FileDownloadLink';
import { SectionCard } from 'Components/Common/SectionCard';
import { ProxyDetails } from 'proto/github.com/solo-io/solo-projects/projects/grpcserver/api/v1/proxy_pb';
import * as React from 'react';
import { useSelector } from 'react-redux';
import { AppState } from 'store';
import { colors, healthConstants } from 'Styles';

const InsideHeader = styled.div`
  display: flex;
  justify-content: space-between;
  font-size: 18px;
  line-height: 22px;
  margin-bottom: 18px;
  color: ${colors.novemberGrey};
`;

const ProxyLogoFullSize = styled(ProxyLogo)`
  width: 33px !important;
  max-height: none !important;
`;

interface Props {}

export const Proxys = (props: Props) => {
  const proxiesList = useSelector(
    (state: AppState) => state.proxies.proxiesList
  );

  const [allProxies, setAllProxies] = React.useState<ProxyDetails.AsObject[]>(
    []
  );

  React.useEffect(() => {
    if (!!proxiesList) {
      const newProxies = proxiesList.filter(proxy => !!proxy.proxy);
      setAllProxies(newProxies);
    }
  }, [proxiesList.length]);

  if (!allProxies.length) {
    return <div>You have no Proxy configurations.</div>;
  }

  return (
    <>
      {allProxies.map((proxy, ind) => {
        return (
          <SectionCard
            key={proxy.proxy!.metadata!.name + ind}
            cardName={proxy.proxy!.metadata!.name}
            logoIcon={<ProxyLogoFullSize />}
            headerSecondaryInformation={[
              {
                title: 'Namespace',
                value: proxy.proxy!.metadata!.namespace
              },
              {
                title: 'Listener Ports',
                value:
                  proxy.proxy!.listenersList.map(l => l.bindPort).join(', ') ||
                  ''
              }
            ]}
            health={
              proxy.proxy!.status
                ? proxy.proxy!.status!.state
                : healthConstants.Pending.value
            }
            healthMessage={'Proxy Status'}>
            <InsideHeader>
              <div>Code Log (Read Only)</div>{' '}
              {!!proxy.raw && (
                <FileDownloadLink
                  fileName={proxy.raw.fileName}
                  fileContent={proxy.raw.content}
                />
              )}
            </InsideHeader>
            {!!proxy.raw && <ConfigDisplayer content={proxy.raw.content} />}
          </SectionCard>
        );
      })}
    </>
  );
};
